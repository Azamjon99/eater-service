package service

import (
	"context"
	"encoding/json"

	"github.com/Azamjon99/eater-service/src/application/dtos"
	ordercanceled "github.com/Azamjon99/eater-service/src/application/integration_events/events/order_cancelled"
	ordercreated "github.com/Azamjon99/eater-service/src/application/integration_events/events/order_created"
	pb "github.com/Azamjon99/eater-service/src/application/protos/eater"
	ordersvc "github.com/Azamjon99/eater-service/src/domain/order/services"
	"github.com/Azamjon99/eater-service/src/infrastructure/pubsub"
	"go.uber.org/zap"
)

type OrderApplicationService interface {
	PlaceOrder(ctx context.Context, request *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error)
	UpdateOrder(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	UpdateOrderByStatus(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	UpdateOrderPaymentByStatus(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	GetOrderById(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	ListOrderByEaterId(ctx context.Context, request *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error)
}

type orderAppSvcImpl struct {
	orderSvc ordersvc.OrderService
    producer pubsub.Producer
    logger *zap.Logger
}

func NewOrderApplicationService(orderSvc ordersvc.OrderService, producer pubsub.Producer,logger *zap.Logger) OrderApplicationService {
	return &orderAppSvcImpl{
		orderSvc: orderSvc,
        producer: producer,
        logger: logger,
	}
}

func (s *orderAppSvcImpl) PlaceOrder(ctx context.Context, request *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
    order, err := s.orderSvc.PlaceOrder(ctx, request.EaterId, request.Instruction, request.RestaurantId)
    if err != nil {
        return nil, err
    }

    orderJson, err := json.Marshal(order)
    if err != nil {
        return nil, err
    }

    var orderMap map[string]interface{}
    if err := json.Unmarshal(orderJson, &orderMap); err != nil {
        return nil, err
    }

    event := ordercreated.NewEvent(order)
    if err:= s.producer.PublishWithKey(event.Topic(), order.ID, event.Payload(), nil); err != nil{
        s.logger.Error("error pushing message", zap.Error(err), zap.String("Topic", event.Topic()))
    }
    response := &pb.PlaceOrderResponse{
        Order: &pb.Order{
        },
    }

    return response, nil
}

func (s *orderAppSvcImpl) UpdateOrder(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.orderSvc.UpdateOrder(ctx, request.Order.Id,  request.Order.Instruction,  request.Order.EntityId)
	if err != nil {
		return nil, err
	}

	orderJson, err := json.Marshal(order)
    if err != nil {
        return nil, err
    }

    var orderMap map[string]interface{}
    if err := json.Unmarshal(orderJson, &orderMap); err != nil {
        return nil, err
    }

    response := &pb.UpdateOrderResponse{
        Order: &pb.Order{
        },
    }

    return response, nil
}

func (s *orderAppSvcImpl) UpdateOrderByStatus(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.orderSvc.UpdateOrderByStatus(ctx, request.Order.Id, request.Order.Status)
	if err != nil {
		return nil, err
	}

	orderJson, err := json.Marshal(order)
    if err != nil {
        return nil, err
    }

    var orderMap map[string]interface{}
    if err := json.Unmarshal(orderJson, &orderMap); err != nil {
        return nil, err
    }
	if request.Order.Status == "2" {
		event := ordercanceled.NewEvent(request.Order.Id)
		if err := s.producer.PublishWithKey(event.Topic(),request.Order.Id,event.Payload(),nil);err!=nil{
			s.logger.Error("error pushing message",zap.Error(err),zap.String("topic",event.Topic()))
		}
	}
    response := &pb.UpdateOrderResponse{
        Order: &pb.Order{
        },
    }

    return response, nil
}

func (s *orderAppSvcImpl) UpdateOrderPaymentByStatus(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	order, err := s.orderSvc.UpdateOrderPaymentByStatus(ctx, request.Order.Id, request.Order.PaymentStatus)
	if err != nil {
		return nil, err
	}

	orderJson, err := json.Marshal(order)
    if err != nil {
        return nil, err
    }

    var orderMap map[string]interface{}
    if err := json.Unmarshal(orderJson, &orderMap); err != nil {
        return nil, err
    }

    response := &pb.UpdateOrderResponse{
        Order: &pb.Order{
        },
    }

    return response, nil
}

func (s *orderAppSvcImpl) DeleteOrder(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	_, err := s.orderSvc.DeleteOrder(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteOrderResponse{}, nil
}

func (s *orderAppSvcImpl) GetOrderById(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	order, err := s.orderSvc.GetOrderById(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}

	orderJson, err := json.Marshal(order)
    if err != nil {
        return nil, err
    }

    var orderMap map[string]interface{}
    if err := json.Unmarshal(orderJson, &orderMap); err != nil {
        return nil, err
    }

    response := &pb.GetOrderResponse{
        Order: &pb.Order{
        },
    }

    return response, nil
}

func (s *orderAppSvcImpl) ListOrderByEaterId(ctx context.Context, request *pb.ListOrderByEaterRequest) (*pb.ListOrderByEaterResponse, error) {
	orders, err := s.orderSvc.ListOrderByEaterId(ctx, request.EaterId, request.Sort, int(request.Page), int(request.PageSize))
	if err != nil {
		return nil, err
	}
        pbOrder := dtos.ConvertToPbOrder(orders)
    
    response := &pb.ListOrderByEaterResponse{
        Orders: pbOrder,
    }

    return response, nil
}

package plandm

type PlanRepository interface {
	Create(*Plan) (*Plan, error)
	FindByID(PlanID) (*Plan, error)
}

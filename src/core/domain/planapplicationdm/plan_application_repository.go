package planapplicationdm

type PlanApplicationRepository interface {
	Create(*PlanApplication) (*PlanApplication, error)
}

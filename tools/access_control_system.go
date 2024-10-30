package tools

// third-party libraries
import "github.com/casbin/casbin"

type Permissions struct {
	// * private * //
	pEnforcer *casbin.Enforcer

	// * public * //
	// the user that wants to access a resource.
	Subject string
	// the resource that is going to be accessed.
	Object string
	// the operation that the user performs on the resource.
	Action string
}

func InitalizeInstance(modelPath string, policyPath string) Permissions {
	// load the access control system
	return Permissions{
		pEnforcer: casbin.NewEnforcer(modelPath, policyPath),
	}
}

func (r *Permissions) Evaluate() bool {
	if res := r.pEnforcer.Enforce(r.Subject, r.Object, r.Action); res {
		return true
	}
	return false
}

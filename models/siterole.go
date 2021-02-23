package models

type SiteRole string

const (
	RoleCreator                   SiteRole = "Creator"
	RoleExplorer                  SiteRole = "Explorer"
	RoleExplorerCanPublish        SiteRole = "ExplorerCanPublish"
	RoleSiteAdministratorExplorer SiteRole = "SiteAdministratorExplorer"
	RoleSiteAdministratorCreator  SiteRole = "SiteAdministratorCreator"
	RoleUnlicensed                SiteRole = "Unlicensed"
	RoleViewer                    SiteRole = "Viewer"
)

func (s *SiteRole) IsValid() bool {
	m := map[SiteRole]bool {
		RoleCreator: true,
		RoleExplorer: true,
		RoleExplorerCanPublish: true,
		RoleSiteAdministratorExplorer: true,
		RoleSiteAdministratorCreator: true,
		RoleUnlicensed: true,
		RoleViewer: true,
	}
	return m[*s]
}

func (s *SiteRole) List() []string {
	return []string {
		string(RoleCreator),
		string(RoleExplorer),
		string(RoleExplorerCanPublish),
		string(RoleSiteAdministratorExplorer),
		string(RoleSiteAdministratorCreator),
		string(RoleUnlicensed),
		string(RoleViewer),
	}
}

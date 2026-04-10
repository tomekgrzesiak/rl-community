// Package domain contains the core domain models for the Spectra Assure Community API.
package domain

// Community represents a supported package repository.
type Community string

const (
	CommunityNPM       Community = "npm"
	CommunityPyPI      Community = "pypi"
	CommunityGem       Community = "gem"
	CommunityNuGet     Community = "nuget"
	CommunityVSX       Community = "vsx"
	CommunityPSGallery Community = "psgallery"
)

// AssessmentStatus represents the result of a policy evaluation.
type AssessmentStatus string

const (
	StatusFail    AssessmentStatus = "fail"
	StatusWarning AssessmentStatus = "warning"
	StatusPass    AssessmentStatus = "pass"
	StatusPending AssessmentStatus = "pending"
)

// SearchRequest represents a single item in a package search request.
type SearchRequest struct {
	UUID            string `json:"uuid"`
	PURL            string `json:"purl,omitempty"`
	SHA1            string `json:"sha1,omitempty"`
	SHA256          string `json:"sha256,omitempty"`
	MatchPattern    string `json:"match_pattern,omitempty"`
	MatchExpression string `json:"match_expression,omitempty"`
	Community       string `json:"community,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Package         string `json:"package,omitempty"`
	Version         string `json:"version,omitempty"`
	Artifact        string `json:"artifact,omitempty"`
	ArtifactTag     string `json:"artifact_tag,omitempty"`
}

// SearchParams holds query parameters for search requests.
type SearchParams struct {
	Offset  int
	Limit   int
	Compact bool
}

// PackageParams holds path and query parameters for package/version endpoints.
type PackageParams struct {
	Community       Community
	Namespace       string
	Package         string
	Version         string
	Artifact        string
	ArtifactTag     string
	MatchPattern    string
	MatchExpression string
	Offset          int
	Limit           int
}

// SearchResponse is the top-level response from the search endpoint.
type SearchResponse struct {
	Community struct {
		Packages []PackageResult `json:"packages"`
		Errors   []SearchError   `json:"errors,omitempty"`
	} `json:"community"`
}

// PackageResult is one result item in a search response.
type PackageResult struct {
	UUID    string      `json:"uuid"`
	Package PackageInfo `json:"package"`
}

// SearchError is an error item in a search response.
type SearchError struct {
	UUID  string `json:"uuid"`
	Error struct {
		Code int    `json:"code"`
		Info string `json:"info"`
	} `json:"error"`
}

// PackageInfo holds metadata and version history for a package.
type PackageInfo struct {
	OwnerVerified  bool            `json:"owner_verified"`
	IsQuarantined  *bool           `json:"is_quarantined"`
	FirstPublished string          `json:"first_published"`
	TotalVersions  int             `json:"total_versions"`
	TotalInstalls  *int            `json:"total_installs"`
	TotalDownloads *int            `json:"total_downloads"`
	LatestVersion  string          `json:"latest_version"`
	DirectImports  *int            `json:"direct_imports"`
	AllMalicious   bool            `json:"all_malicious"`
	WasArchived    *bool           `json:"was_archived"`
	WasRemoved     bool            `json:"was_removed"`
	KeyProject     bool            `json:"key_project"`
	Popularity     *int            `json:"popularity"`
	Identity       Identity        `json:"identity"`
	Versions       []VersionSummary `json:"versions"`
	Incidents      IncidentStats   `json:"incidents"`
	Contributors   []Contributor   `json:"contributors,omitempty"`
}

// Identity is the PURL-based identification of a package.
type Identity struct {
	PURL        string    `json:"purl"`
	Community   Community `json:"community"`
	Namespace   string    `json:"namespace"`
	Package     string    `json:"package"`
	Product     string    `json:"product"`
	Version     *string   `json:"version"`
	Artifact    *string   `json:"artifact"`
	License     string    `json:"license"`
	Published   string    `json:"published"`
	Deprecated  *bool     `json:"deprecated"`
	Removed     bool      `json:"removed"`
	Category    *string   `json:"category"`
	Homepage    string    `json:"homepage"`
	Repository  string    `json:"repository"`
	Description string    `json:"description"`
	Keywords    []string  `json:"keywords,omitempty"`
}

// VersionSummary is a brief entry in a package's version list.
type VersionSummary struct {
	Version   string          `json:"version"`
	Published string          `json:"published"`
	Quality   *QualitySummary `json:"quality,omitempty"`
	Artifacts []Artifact      `json:"artifacts,omitempty"`
	Assessments *Assessments  `json:"assessments,omitempty"`
	Incidents map[string]Incident `json:"incidents,omitempty"`
}

// QualitySummary is the condensed quality rating for a version.
type QualitySummary struct {
	Status     AssessmentStatus `json:"status"`
	Priority   *int             `json:"priority"`
	Assessment string           `json:"assessment"`
	Metrics    *QualityMetrics  `json:"metrics,omitempty"`
}

// QualityMetrics holds policy violation counts.
type QualityMetrics struct {
	High   int `json:"high"`
	Medium int `json:"medium"`
	Low    int `json:"low"`
}

// Artifact describes a linked or generated artifact file.
type Artifact struct {
	Type string `json:"type"`
	Ref  string `json:"ref"`
}

// Assessments holds the per-category risk assessments.
type Assessments struct {
	Licenses        *Assessment `json:"licenses,omitempty"`
	Malware         *Assessment `json:"malware,omitempty"`
	Hardening       *Assessment `json:"hardening,omitempty"`
	Secrets         *Assessment `json:"secrets,omitempty"`
	Tampering       *Assessment `json:"tampering,omitempty"`
	Vulnerabilities *Assessment `json:"vulnerabilities,omitempty"`
	Repository      *Assessment `json:"repository,omitempty"`
}

// Assessment is a single risk assessment result.
type Assessment struct {
	Final      bool             `json:"final"`
	Enabled    bool             `json:"enabled"`
	Priority   *int             `json:"priority"`
	Label      string           `json:"label"`
	Violations []string         `json:"violations"`
	Count      int              `json:"count"`
	Status     AssessmentStatus `json:"status"`
}

// Incident describes a malware or removal incident.
type Incident struct {
	Type      string     `json:"type"`
	Reporters []Reporter `json:"reporters"`
}

// Reporter is a source for an incident report.
type Reporter struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Date string `json:"date"`
	Link string `json:"link"`
}

// IncidentStats is the aggregate incident statistics for a package.
type IncidentStats struct {
	Malware      int            `json:"malware"`
	RecentMalware RecentIncident `json:"recent_malware"`
	Removal      int            `json:"removal"`
	RecentRemoval RecentIncident `json:"recent_removal"`
}

// RecentIncident summarises the most recent incident of a given type.
type RecentIncident struct {
	LatestVersion   string `json:"latest_version"`
	LatestTimestamp string `json:"latest_timestamp"`
	RecentCount     int    `json:"recent_count"`
}

// Contributor is a person associated with a package project.
type Contributor struct {
	Role  string `json:"role"`
	Name  string `json:"name"`
	User  string `json:"user"`
	Email string `json:"email"`
}

// VersionReportResponse is the top-level response from the version report endpoint.
type VersionReportResponse struct {
	Community struct {
		Report VersionReport `json:"report"`
	} `json:"community"`
}

// VersionReport is the detailed analysis report for one package version.
type VersionReport struct {
	Info     ReportInfo     `json:"info"`
	Metadata *ReportMetadata `json:"metadata,omitempty"`
}

// ReportInfo is the summary info block within a version report.
type ReportInfo struct {
	File     FileInfo     `json:"file"`
	Analysis AnalysisInfo `json:"analysis"`
	Statistics *Statistics `json:"statistics,omitempty"`
	Detections map[string]map[string]int `json:"detections,omitempty"`
	Disabled  []string    `json:"disabled,omitempty"`
}

// FileInfo describes the analysed package file.
type FileInfo struct {
	Name      string   `json:"name"`
	Size      *int     `json:"size"`
	SourceURL string   `json:"source_url"`
	Downloaded bool    `json:"downloaded"`
	Hashes    [][]string `json:"hashes"`
	Identity  *Identity `json:"identity,omitempty"`
}

// AnalysisInfo holds the analysis engine metadata.
type AnalysisInfo struct {
	Version   string  `json:"version"`
	Catalogue *int    `json:"catalogue"`
	Timestamp string  `json:"timestamp"`
}

// Statistics contains aggregated analysis statistics.
type Statistics struct {
	Components    int              `json:"components"`
	Dependencies  int              `json:"dependencies"`
	Extracted     int              `json:"extracted"`
	License       *LicenseStats    `json:"license,omitempty"`
	Vulnerabilities *VulnStats     `json:"vulnerabilities,omitempty"`
	Quality       *FullQuality     `json:"quality,omitempty"`
}

// LicenseStats is the license type breakdown.
type LicenseStats struct {
	Undeclared   int `json:"undeclared"`
	PublicDomain int `json:"public_domain"`
	Permissive   int `json:"permissive"`
	WeakCopyleft int `json:"weak_copyleft"`
	Copyleft     int `json:"copyleft"`
	Freeware     int `json:"freeware"`
	Shareware    int `json:"shareware"`
	Freemium     int `json:"freemium"`
	NonCommercial int `json:"non_commercial"`
	Proprietary  int `json:"proprietary"`
}

// VulnStats is the vulnerability severity breakdown.
type VulnStats struct {
	Total    int `json:"total"`
	Critical int `json:"critical"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
	Exploit  int `json:"exploit"`
	Malware  int `json:"malware"`
	Mandate  int `json:"mandate"`
	Fixable  int `json:"fixable"`
	Named    int `json:"named"`
}

// FullQuality is the detailed policy evaluation result.
type FullQuality struct {
	Status   string       `json:"status"`
	Priority int          `json:"priority"`
	Metrics  PolicyMetrics `json:"metrics"`
}

// PolicyMetrics holds the full policy check counters.
type PolicyMetrics struct {
	Total   int `json:"total"`
	Pass    int `json:"pass"`
	Warning int `json:"warning"`
	Fail    int `json:"fail"`
	High    int `json:"high"`
	Medium  int `json:"medium"`
	Low     int `json:"low"`
}

// ReportMetadata is the detailed metadata block in a version report.
type ReportMetadata struct {
	Assessments     *Assessments               `json:"assessments,omitempty"`
	Incidents       map[string]Incident        `json:"incidents,omitempty"`
	Violations      map[string]Violation       `json:"violations,omitempty"`
	Indicators      map[string]Indicator       `json:"indicators,omitempty"`
	Classifications []Classification           `json:"classifications,omitempty"`
	Vulnerabilities map[string]Vulnerability   `json:"vulnerabilities,omitempty"`
	Dependencies    map[string]Dependency      `json:"dependencies,omitempty"`
}

// Violation is a policy check result.
type Violation struct {
	RuleID      string     `json:"rule_id"`
	Status      string     `json:"status"`
	Priority    int        `json:"priority"`
	Severity    string     `json:"severity"`
	Category    string     `json:"category"`
	Description string     `json:"description"`
	Statistics  ViolationStats `json:"statistics"`
}

// ViolationStats is the per-file enforcement stats for a violation.
type ViolationStats struct {
	Applicable   int `json:"applicable"`
	Enforcements int `json:"enforcements"`
	Exclusions   int `json:"exclusions"`
	Violations   int `json:"violations"`
}

// Indicator is a behavior indicator entry.
type Indicator struct {
	RuleID      string `json:"rule_id"`
	Priority    int    `json:"priority"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Exclusions  int    `json:"exclusions"`
	Occurrences int    `json:"occurrences"`
	Violations  int    `json:"violations"`
}

// Classification is a file threat classification result.
type Classification struct {
	Object string     `json:"object"`
	Status string     `json:"status"`
	Result string     `json:"result"`
	Hashes [][]string `json:"hashes"`
}

// Vulnerability is a CVE entry with CVSS and exploit details.
type Vulnerability struct {
	Name    string      `json:"name"`
	CVSS    CVSSInfo    `json:"cvss"`
	Updated string      `json:"updated"`
	Affects []string    `json:"affects"`
	Summary string      `json:"summary"`
	Audit   AuditInfo   `json:"audit"`
	Exploit []string    `json:"exploit"`
}

// CVSSInfo holds CVSS version and score.
type CVSSInfo struct {
	Version   string  `json:"version"`
	BaseScore float64 `json:"baseScore"`
}

// AuditInfo holds triage audit trail info for a vulnerability.
type AuditInfo struct {
	Author    string `json:"author"`
	Timestamp string `json:"timestamp"`
	Reason    string `json:"reason"`
}

// Dependency is a declared package dependency.
type Dependency struct {
	Type            string  `json:"type"`
	PURL            string  `json:"purl"`
	Community       string  `json:"community"`
	Framework       *string `json:"framework"`
	Product         *string `json:"product"`
	Version         string  `json:"version"`
	License         *string `json:"license"`
	Vulnerabilities []string `json:"vulnerabilities"`
	Classification  DependencyClassification `json:"classification"`
}

// DependencyClassification classifies a dependency as malicious/suspicious.
type DependencyClassification struct {
	Status *string `json:"status"`
	Result string  `json:"result"`
}

// PackageDetailsResponse is the top-level response from the package details endpoint.
type PackageDetailsResponse struct {
	Community struct {
		Package PackageInfo `json:"package"`
	} `json:"community"`
}

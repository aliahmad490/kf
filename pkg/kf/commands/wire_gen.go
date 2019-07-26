// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package commands

import (
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/kf/pkg/client/clientset/versioned/typed/kf/v1alpha1"
	"github.com/google/kf/pkg/kf"
	"github.com/google/kf/pkg/kf/apps"
	"github.com/google/kf/pkg/kf/buildpacks"
	builds2 "github.com/google/kf/pkg/kf/builds"
	apps2 "github.com/google/kf/pkg/kf/commands/apps"
	buildpacks2 "github.com/google/kf/pkg/kf/commands/buildpacks"
	"github.com/google/kf/pkg/kf/commands/builds"
	"github.com/google/kf/pkg/kf/commands/config"
	"github.com/google/kf/pkg/kf/commands/quotas"
	routes2 "github.com/google/kf/pkg/kf/commands/routes"
	servicebindings2 "github.com/google/kf/pkg/kf/commands/service-bindings"
	services2 "github.com/google/kf/pkg/kf/commands/services"
	spaces2 "github.com/google/kf/pkg/kf/commands/spaces"
	"github.com/google/kf/pkg/kf/logs"
	"github.com/google/kf/pkg/kf/routes"
	"github.com/google/kf/pkg/kf/service-bindings"
	"github.com/google/kf/pkg/kf/services"
	"github.com/google/kf/pkg/kf/sources"
	"github.com/google/kf/pkg/kf/spaces"
	"github.com/google/kf/pkg/kf/systemenvinjector"
	"github.com/google/wire"
	logs2 "github.com/knative/build/pkg/logs"
	"github.com/poy/kontext"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

import (
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

// Injectors from wire_injector.go:

func InjectPush(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	pusher := apps.NewPusher(appsClient)
	srcImageBuilder := provideSrcImageBuilder()
	servicecatalogV1beta1Interface := config.GetServiceCatalogClient(p)
	clientInterface := config.GetSecretClient(p)
	servicebindingsClientInterface := servicebindings.NewClient(servicecatalogV1beta1Interface, clientInterface)
	command := apps2.NewPushCommand(p, appsClient, pusher, srcImageBuilder, servicebindingsClientInterface)
	return command
}

func InjectDelete(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewDeleteCommand(p, appsClient)
	return command
}

func InjectApps(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewAppsCommand(p, appsClient)
	return command
}

func InjectGetApp(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewGetAppCommand(p, appsClient)
	return command
}

func InjectScale(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewScaleCommand(p, appsClient)
	return command
}

func InjectStart(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewStartCommand(p, appsClient)
	return command
}

func InjectStop(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewStopCommand(p, appsClient)
	return command
}

func InjectRestart(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewRestartCommand(p, appsClient)
	return command
}

func InjectRestage(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewRestageCommand(p, appsClient)
	return command
}

func InjectProxy(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	kubernetesInterface := config.GetKubernetes(p)
	ingressLister := kf.NewIstioClient(kubernetesInterface)
	command := apps2.NewProxyCommand(p, appsClient, ingressLister)
	return command
}

func InjectLogs(p *config.KfParams) *cobra.Command {
	coreV1Interface := provideCoreV1(p)
	tailer := logs.NewTailer(coreV1Interface)
	command := apps2.NewLogsCommand(p, tailer)
	return command
}

func InjectEnv(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewEnvCommand(p, appsClient)
	return command
}

func InjectSetEnv(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewSetEnvCommand(p, appsClient)
	return command
}

func InjectUnsetEnv(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, client)
	command := apps2.NewUnsetEnvCommand(p, appsClient)
	return command
}

func provideSystemEnvInjector(p *config.KfParams) systemenvinjector.SystemEnvInjectorInterface {
	servicecatalogV1beta1Interface := config.GetServiceCatalogClient(p)
	clientInterface := config.GetSecretClient(p)
	servicebindingsClientInterface := servicebindings.NewClient(servicecatalogV1beta1Interface, clientInterface)
	systemEnvInjectorInterface := systemenvinjector.NewSystemEnvInjector(servicebindingsClientInterface)
	return systemEnvInjectorInterface
}

func InjectCreateService(p *config.KfParams) *cobra.Command {
	sClientFactory := config.GetSvcatApp(p)
	clientInterface := services.NewClient(sClientFactory)
	command := services2.NewCreateServiceCommand(p, clientInterface)
	return command
}

func InjectDeleteService(p *config.KfParams) *cobra.Command {
	sClientFactory := config.GetSvcatApp(p)
	clientInterface := services.NewClient(sClientFactory)
	command := services2.NewDeleteServiceCommand(p, clientInterface)
	return command
}

func InjectGetService(p *config.KfParams) *cobra.Command {
	sClientFactory := config.GetSvcatApp(p)
	clientInterface := services.NewClient(sClientFactory)
	command := services2.NewGetServiceCommand(p, clientInterface)
	return command
}

func InjectListServices(p *config.KfParams) *cobra.Command {
	sClientFactory := config.GetSvcatApp(p)
	clientInterface := services.NewClient(sClientFactory)
	command := services2.NewListServicesCommand(p, clientInterface)
	return command
}

func InjectMarketplace(p *config.KfParams) *cobra.Command {
	sClientFactory := config.GetSvcatApp(p)
	clientInterface := services.NewClient(sClientFactory)
	command := services2.NewMarketplaceCommand(p, clientInterface)
	return command
}

func InjectBindingService(p *config.KfParams) *cobra.Command {
	servicecatalogV1beta1Interface := config.GetServiceCatalogClient(p)
	clientInterface := config.GetSecretClient(p)
	servicebindingsClientInterface := servicebindings.NewClient(servicecatalogV1beta1Interface, clientInterface)
	command := servicebindings2.NewBindServiceCommand(p, servicebindingsClientInterface)
	return command
}

func InjectListBindings(p *config.KfParams) *cobra.Command {
	servicecatalogV1beta1Interface := config.GetServiceCatalogClient(p)
	clientInterface := config.GetSecretClient(p)
	servicebindingsClientInterface := servicebindings.NewClient(servicecatalogV1beta1Interface, clientInterface)
	command := servicebindings2.NewListBindingsCommand(p, servicebindingsClientInterface)
	return command
}

func InjectUnbindService(p *config.KfParams) *cobra.Command {
	servicecatalogV1beta1Interface := config.GetServiceCatalogClient(p)
	clientInterface := config.GetSecretClient(p)
	servicebindingsClientInterface := servicebindings.NewClient(servicecatalogV1beta1Interface, clientInterface)
	command := servicebindings2.NewUnbindServiceCommand(p, servicebindingsClientInterface)
	return command
}

func InjectVcapServices(p *config.KfParams) *cobra.Command {
	servicecatalogV1beta1Interface := config.GetServiceCatalogClient(p)
	clientInterface := config.GetSecretClient(p)
	servicebindingsClientInterface := servicebindings.NewClient(servicecatalogV1beta1Interface, clientInterface)
	command := servicebindings2.NewVcapServicesCommand(p, servicebindingsClientInterface)
	return command
}

func InjectBuildpacksClient(p *config.KfParams) buildpacks.Client {
	remoteImageFetcher := provideRemoteImageFetcher()
	client := buildpacks.NewClient(remoteImageFetcher)
	return client
}

func InjectBuildpacks(p *config.KfParams) *cobra.Command {
	client := InjectBuildpacksClient(p)
	command := buildpacks2.NewBuildpacksCommand(p, client)
	return command
}

func InjectStacks(p *config.KfParams) *cobra.Command {
	client := InjectBuildpacksClient(p)
	command := buildpacks2.NewStacksCommand(p, client)
	return command
}

func InjectSpaces(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := spaces2.NewListSpacesCommand(p, client)
	return command
}

func InjectSpace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := spaces2.NewGetSpaceCommand(p, client)
	return command
}

func InjectCreateSpace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := spaces2.NewCreateSpaceCommand(p, client)
	return command
}

func InjectDeleteSpace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := spaces2.NewDeleteSpaceCommand(p, client)
	return command
}

func InjectConfigSpace(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := spaces2.NewConfigSpaceCommand(p, client)
	return command
}

func InjectCreateQuota(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := quotas.NewCreateQuotaCommand(p, client)
	return command
}

func InjectUpdateQuota(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := quotas.NewUpdateQuotaCommand(p, client)
	return command
}

func InjectGetQuota(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := quotas.NewGetQuotaCommand(p, client)
	return command
}

func InjectDeleteQuota(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	spacesGetter := provideKfSpaces(kfV1alpha1Interface)
	client := spaces.NewClient(spacesGetter)
	command := quotas.NewDeleteQuotaCommand(p, client)
	return command
}

func InjectRoutes(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes.NewClient(kfV1alpha1Interface)
	command := routes2.NewRoutesCommand(p, client)
	return command
}

func InjectCreateRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes.NewClient(kfV1alpha1Interface)
	command := routes2.NewCreateRouteCommand(p, client)
	return command
}

func InjectDeleteRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes.NewClient(kfV1alpha1Interface)
	command := routes2.NewDeleteRouteCommand(p, client)
	return command
}

func InjectMapRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes.NewClient(kfV1alpha1Interface)
	appsGetter := provideAppsGetter(kfV1alpha1Interface)
	systemEnvInjectorInterface := provideSystemEnvInjector(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	sourcesClient := sources.NewClient(sourcesGetter, buildTailer)
	appsClient := apps.NewClient(appsGetter, systemEnvInjectorInterface, sourcesClient)
	command := routes2.NewMapRouteCommand(p, client, appsClient)
	return command
}

func InjectUnmapRoute(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	client := routes.NewClient(kfV1alpha1Interface)
	command := routes2.NewUnmapRouteCommand(p, client)
	return command
}

func InjectBuilds(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	command := builds.NewListBuildsCommand(p, client)
	return command
}

func InjectBuildLogs(p *config.KfParams) *cobra.Command {
	kfV1alpha1Interface := config.GetKfClient(p)
	sourcesGetter := provideKfSources(kfV1alpha1Interface)
	buildTailer := provideSourcesBuildTailer()
	client := sources.NewClient(sourcesGetter, buildTailer)
	command := builds.NewBuildLogsCommand(p, client)
	return command
}

// wire_injector.go:

func provideSrcImageBuilder() apps2.SrcImageBuilder {
	return apps2.SrcImageBuilderFunc(kontext.BuildImage)
}

func provideBuildTailer() builds2.BuildTailer {
	return builds2.BuildTailerFunc(logs2.Tail)
}

var AppsSet = wire.NewSet(
	SourcesSet,
	provideAppsGetter, apps.NewClient, apps.NewPusher, provideSystemEnvInjector,
)

func provideAppsGetter(ki v1alpha1.KfV1alpha1Interface) v1alpha1.AppsGetter {
	return ki
}

func provideCoreV1(p *config.KfParams) v1.CoreV1Interface {
	return config.GetKubernetes(p).CoreV1()
}

/////////////////
// Buildpacks //
///////////////
func provideRemoteImageFetcher() buildpacks.RemoteImageFetcher {
	return remote.Image
}

var SpacesSet = wire.NewSet(config.GetKfClient, provideKfSpaces, spaces.NewClient)

func provideKfSpaces(ki v1alpha1.KfV1alpha1Interface) v1alpha1.SpacesGetter {
	return ki
}

var SourcesSet = wire.NewSet(config.GetKfClient, provideSourcesBuildTailer, provideKfSources, sources.NewClient)

func provideKfSources(ki v1alpha1.KfV1alpha1Interface) v1alpha1.SourcesGetter {
	return ki
}

func provideSourcesBuildTailer() sources.BuildTailer {
	return builds2.BuildTailerFunc(logs2.Tail)
}

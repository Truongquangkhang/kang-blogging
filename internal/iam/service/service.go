package service

// import (
// 	"context"

// 	"github.com/sirupsen/logrus"
// )

// func NewApplication(ctx context.Context) (app.Application, func()) {
// }

// func newApplication(
// 	ctx context.Context,
// 	nioCoreClient applicable_vouchers.ApplicableVouchersServiceClient,
// 	qrsearchClient qrsearch.ClientWithResponses,
// 	mktClient mktportal.ClientWithResponses,
// ) app.Application {
// 	logrus.Info(ctx)

// 	nioCoreService := niocore.NewNioCoreGrpc(nioCoreClient)
// 	mktPortalService := mktportalsvc.NewMktPortalHttp(mktClient)
// 	qrsearchsvc := qrsearchsvc.NewQrSearchHttp(qrsearchClient)
// 	voucherRepository := adapters.NewVoucherRepositoryImpl(
// 		nioCoreService,
// 		qrsearchsvc,
// 		mktPortalService,
// 	)
// 	bannerRepository := adapters.NewBannerRepositoryImpl()

// 	logger := logrus.NewEntry(logrus.StandardLogger())

// 	metricsClient := metrics.NoOp{}

// 	return app.Application{
// 		Usecases: app.Usecases{
// 			GetVouchersForShop:  usecase.NewGetVoucherForShopsHandler(voucherRepository, logger, metricsClient),
// 			GetAvailableVoucher: usecase.NewGetAvailableVoucherHandler(voucherRepository, logger, metricsClient),
// 			GetBreakingNews:     usecase.NewGetBreakingNewsHandler(bannerRepository, logger, metricsClient),
// 			ExtractQr:           usecase.NewExtractQrHandler(logger, metricsClient),
// 			UseVoucher:          usecase.NewUseVoucherHandler(voucherRepository, logger, metricsClient),
// 		},
// 	}
// }

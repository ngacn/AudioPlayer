#include <QApplication>

#include "mainwindow.h"

int main(int argc, char *argv[])
{
    Q_INIT_RESOURCE(application);

    QApplication app(argc, argv);
    app.setOrganizationName("ngacn@prognyan");
    app.setApplicationName("Warpten Player");
    MainWindow mainWin;
    mainWin.show();
    return app.exec();
}

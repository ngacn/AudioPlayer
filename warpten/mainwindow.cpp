#include <QtWidgets>

#include "mainwindow.h"
#include "playlisttab.h"
#include "warptencli.h"

MainWindow::MainWindow()
{
    daemonProcess = new QProcess(this);
    QStringList args;
    args << "-d" << "-t";
    daemonProcess->start("warpten-daemon", args);

    requestVersion();
    requestPlaylists();

    playlistsTabWidget = new QTabWidget;
    playlistsTabWidget->setTabsClosable(true);
    setCentralWidget(playlistsTabWidget);
    connect(playlistsTabWidget, SIGNAL(tabCloseRequested(int)), this, SLOT(requestCloseTab(int)));

    createActions();
    createMenus();
    createStatusBar();

    readSettings();

    setUnifiedTitleAndToolBarOnMac(true);
}

void MainWindow::closeEvent(QCloseEvent *event)
{
    daemonProcess->terminate();
    writeSettings();
    event->accept();
}

void MainWindow::newTrack()
{
    QStringList files = QFileDialog::getOpenFileNames(
                            this,
                            "Select one or more files to add",
                            QDir::homePath());
    PlaylistTab *curTab = static_cast<PlaylistTab*>(playlistsTabWidget->currentWidget());
    QString curUuid = curTab->getUuid();
    QStringList::Iterator it = files.begin();
    while(it != files.end()) {
        QUrlQuery query;
        query.addQueryItem("path", *it);
        query.addQueryItem("playlist", curUuid);
        WarptenCli *cli = new WarptenCli(this);
        connect(cli, SIGNAL(on_execution_finished(WarptenCli*)), this, SLOT(updateNewTrack(WarptenCli*)));
        cli->execute("POST", "/track/add", query);

        ++it;
    }
}

void MainWindow::newPlaylist()
{
    bool ok;
    QString text = QInputDialog::getText(this, tr("InputDialog"),
                                         tr("New Playlist:"), QLineEdit::Normal,
                                         NULL, &ok);
    if (ok && !text.isEmpty()) {
        QUrlQuery query;
        query.addQueryItem("name", text);
        WarptenCli *cli = new WarptenCli(this);
        connect(cli, SIGNAL(on_execution_finished(WarptenCli*)), this, SLOT(updateNewPlaylist(WarptenCli*)));
        cli->execute("POST", "/playlist/add", query);
    }
}

void MainWindow::about()
{
    QMessageBox::about(this, tr("About Warpten"),
                       tr("<b>Warpten</b> v%1 is an audio player similar to foobar2000.").arg(version));
}

void MainWindow::createActions()
{
    newTrackAct = new QAction(tr("&Add file(s)"), this);
    newTrackAct->setShortcuts(QKeySequence::Open);
    newTrackAct->setStatusTip(tr("Add file(s) to current playlist"));
    connect(newTrackAct, SIGNAL(triggered()), this, SLOT(newTrack()));

    newPlaylistAct = new QAction(QIcon(":/images/new.png"), tr("&New playlist"), this);
    newPlaylistAct->setShortcuts(QKeySequence::AddTab);
    newPlaylistAct->setStatusTip(tr("Create a new playlist"));
    connect(newPlaylistAct, SIGNAL(triggered()), this, SLOT(newPlaylist()));

    exitAct = new QAction(tr("E&xit"), this);
    exitAct->setShortcuts(QKeySequence::Quit);
    exitAct->setStatusTip(tr("Exit Warpten"));
    connect(exitAct, SIGNAL(triggered()), this, SLOT(close()));

    aboutAct = new QAction(tr("&About"), this);
    aboutAct->setStatusTip(tr("Show Warpten's About box"));
    connect(aboutAct, SIGNAL(triggered()), this, SLOT(about()));

    aboutQtAct = new QAction(tr("About &Qt"), this);
    aboutQtAct->setStatusTip(tr("Show the Qt library's About box"));
    connect(aboutQtAct, SIGNAL(triggered()), qApp, SLOT(aboutQt()));
}

void MainWindow::createMenus()
{
    fileMenu = menuBar()->addMenu(tr("&File"));
    fileMenu->addAction(newTrackAct);
    fileMenu->addSeparator();
    fileMenu->addAction(newPlaylistAct);
    fileMenu->addSeparator();
    fileMenu->addAction(exitAct);

    editMenu = menuBar()->addMenu(tr("&Edit"));

    menuBar()->addSeparator();

    helpMenu = menuBar()->addMenu(tr("&Help"));
    helpMenu->addAction(aboutAct);
    helpMenu->addAction(aboutQtAct);
}

void MainWindow::createStatusBar()
{
    statusBar()->showMessage(tr("Ready"));
}

void MainWindow::readSettings()
{
    QSettings settings("Warpten", "WarptenPlayer");
    QPoint pos = settings.value("pos", QPoint(200, 200)).toPoint();
    QSize size = settings.value("size", QSize(400, 400)).toSize();
    resize(size);
    move(pos);
}

void MainWindow::writeSettings()
{
    QSettings settings("Warpten", "WarptenPlayer");
    settings.setValue("pos", pos());
    settings.setValue("size", size());
}

void MainWindow::requestVersion()
{
    QUrlQuery query;
    WarptenCli *cli = new WarptenCli(this);
    connect(cli, SIGNAL(on_execution_finished(WarptenCli*)), this, SLOT(updateVersion(WarptenCli*)));
    cli->execute("GET", "/version", query);
}

void MainWindow::updateVersion(WarptenCli *cli)
{
    if (cli->errorType != QNetworkReply::NoError) {
        // an error occurred
        return;
    }
    QJsonDocument loadDoc(QJsonDocument::fromJson(cli->response));
    QJsonObject json = loadDoc.object();
    if (!json["err"].toString().isEmpty()) {
        qDebug() << json["err"].toString();
        return;
    }
    version = json["return"].toString();
}

void MainWindow::requestPlaylists()
{
    QUrlQuery query;
    WarptenCli *cli = new WarptenCli(this);
    connect(cli, SIGNAL(on_execution_finished(WarptenCli*)), this, SLOT(updatePlaylists(WarptenCli*)));
    cli->execute("GET", "/playlists", query);
}

void MainWindow::updatePlaylists(WarptenCli *cli)
{
    if (cli->errorType != QNetworkReply::NoError) {
        // an error occurred
        return;
    }
    QJsonDocument loadDoc(QJsonDocument::fromJson(cli->response));
    QJsonObject json = loadDoc.object();
    if (!json["err"].toString().isEmpty()) {
        qDebug() << json["err"].toString();
        return;
    }
    QJsonObject pls = json["return"].toObject();
    foreach (const QJsonValue pl, pls) {
        QString name = pl.toObject()["name"].toString();
        QString uuid = pl.toObject()["uuid"].toString();
        playlistsTabWidget->addTab(new PlaylistTab(uuid), name);
    }
}

void MainWindow::updateNewTrack(WarptenCli *cli)
{
    if (cli->errorType != QNetworkReply::NoError) {
        // an error occurred
        return;
    }
    QJsonDocument loadDoc(QJsonDocument::fromJson(cli->response));
    QJsonObject json = loadDoc.object();
    if (!json["err"].toString().isEmpty()) {
        qDebug() << json["err"].toString();
        return;
    }
    PlaylistTab *curTab = static_cast<PlaylistTab*>(playlistsTabWidget->currentWidget());
    QJsonObject tk = json["return"].toObject();
    QString path = tk["path"].toString();
    QString uuid = tk["uuid"].toString();
    curTab->addTrack(uuid, path);
}

void MainWindow::updateNewPlaylist(WarptenCli *cli)
{
    if (cli->errorType != QNetworkReply::NoError) {
        // an error occurred
        return;
    }
    QJsonDocument loadDoc(QJsonDocument::fromJson(cli->response));
    QJsonObject json = loadDoc.object();
    if (!json["err"].toString().isEmpty()) {
        qDebug() << json["err"].toString();
        return;
    }
    QJsonObject pl = json["return"].toObject();
    QString name = pl["name"].toString();
    QString uuid = pl["uuid"].toString();
    playlistsTabWidget->addTab(new PlaylistTab(uuid), name);
}

void MainWindow::requestCloseTab(int index)
{
    PlaylistTab *tab = static_cast<PlaylistTab*>(playlistsTabWidget->widget(index));
    QUrlQuery query;
    query.addQueryItem("uuid", tab->getUuid());
    query.addQueryItem("index", QString::number(index));
    WarptenCli *cli = new WarptenCli(this);
    connect(cli, SIGNAL(on_execution_finished(WarptenCli*)), this, SLOT(updateCloseTab(WarptenCli*)));
    cli->execute("POST", "/playlist/del", query);
}

void MainWindow::updateCloseTab(WarptenCli *cli)
{
    if (cli->errorType != QNetworkReply::NoError) {
        // an error occurred
        return;
    }
    QJsonDocument loadDoc(QJsonDocument::fromJson(cli->response));
    QJsonObject json = loadDoc.object();
    if (!json["err"].toString().isEmpty()) {
        qDebug() << json["err"].toString();
        return;
    }
    int index = json["return"].toInt();
    delete playlistsTabWidget->widget(index);
}

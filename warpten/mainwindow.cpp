#include <QtWidgets>

#include "mainwindow.h"
#include "playlisttab.h"

MainWindow::MainWindow()
{
    playlistsTabWidget = new QTabWidget;
    playlistsTabWidget->addTab(new PlaylistTab(), tr("Default"));
    setCentralWidget(playlistsTabWidget);

    createActions();
    createMenus();
    createStatusBar();

    readSettings();

    setUnifiedTitleAndToolBarOnMac(true);
}

void MainWindow::closeEvent(QCloseEvent *event)
{
    writeSettings();
    event->accept();
}

void MainWindow::newPlaylist()
{
    playlistsTabWidget->addTab(new PlaylistTab(), tr("NewList"));
}

void MainWindow::about()
{
   QMessageBox::about(this, tr("About Warpten"),
            tr("<b>Warpten</b> is an audio player similar to foobar2000."));
}

void MainWindow::createActions()
{
    newPlaylistAct = new QAction(QIcon(":/images/new.png"), tr("&New playlist"), this);
    newPlaylistAct->setShortcuts(QKeySequence::New);
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
    QSettings settings("Warpten", "Warpten Player");
    QPoint pos = settings.value("pos", QPoint(200, 200)).toPoint();
    QSize size = settings.value("size", QSize(400, 400)).toSize();
    resize(size);
    move(pos);
}

void MainWindow::writeSettings()
{
    QSettings settings("Warpten", "Warpten Player");
    settings.setValue("pos", pos());
    settings.setValue("size", size());
}

#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QProcess>

QT_BEGIN_NAMESPACE
class QAction;
class QMenu;
QT_END_NAMESPACE

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow();

protected:
    void closeEvent(QCloseEvent *event);

private slots:
    void newPlaylist();
    void about();

private:
    void createActions();
    void createMenus();
    void createStatusBar();
    void readSettings();
    void writeSettings();

    QMenu *fileMenu;
    QMenu *editMenu;
    QMenu *helpMenu;
    QAction *newPlaylistAct;
    QAction *exitAct;
    QAction *aboutAct;
    QAction *aboutQtAct;
    QTabWidget *playlistsTabWidget;

    QProcess *daemonProcess;
};

#endif

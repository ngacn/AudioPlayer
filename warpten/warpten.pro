QT += widgets

HEADERS       = mainwindow.h \
    playlisttab.h
SOURCES       = main.cpp \
                mainwindow.cpp \
    playlisttab.cpp
RESOURCES     = warpten.qrc

# install
target.path = $$[QT_INSTALL_EXAMPLES]/widgets/mainwindows/warpten
INSTALLS += target

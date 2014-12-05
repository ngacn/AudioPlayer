QT += widgets network

HEADERS       = mainwindow.h \
    playlisttab.h \
    warptencli.h
SOURCES       = main.cpp \
                mainwindow.cpp \
    playlisttab.cpp \
    warptencli.cpp
RESOURCES     = warpten.qrc

# install
target.path = $$[QT_INSTALL_EXAMPLES]/widgets/mainwindows/warpten
INSTALLS += target

package main

// XorgServer - X Window System display server
// Homepage: https://www.x.org

import (
	"fmt"
	
	"os/exec"
)

func installXorgServer() {
	// Método 1: Descargar y extraer .tar.gz
	xorgserver_tar_url := "https://www.x.org/releases/individual/xserver/xorg-server-21.1.13.tar.xz"
	xorgserver_cmd_tar := exec.Command("curl", "-L", xorgserver_tar_url, "-o", "package.tar.gz")
	err := xorgserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xorgserver_zip_url := "https://www.x.org/releases/individual/xserver/xorg-server-21.1.13.tar.xz"
	xorgserver_cmd_zip := exec.Command("curl", "-L", xorgserver_zip_url, "-o", "package.zip")
	err = xorgserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xorgserver_bin_url := "https://www.x.org/releases/individual/xserver/xorg-server-21.1.13.tar.xz"
	xorgserver_cmd_bin := exec.Command("curl", "-L", xorgserver_bin_url, "-o", "binary.bin")
	err = xorgserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xorgserver_src_url := "https://www.x.org/releases/individual/xserver/xorg-server-21.1.13.tar.xz"
	xorgserver_cmd_src := exec.Command("curl", "-L", xorgserver_src_url, "-o", "source.tar.gz")
	err = xorgserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xorgserver_cmd_direct := exec.Command("./binary")
	err = xorgserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: font-util")
exec.Command("latte", "install", "font-util")
	fmt.Println("Instalando dependencia: libxkbfile")
exec.Command("latte", "install", "libxkbfile")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
	fmt.Println("Instalando dependencia: xtrans")
exec.Command("latte", "install", "xtrans")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxau")
exec.Command("latte", "install", "libxau")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: libxdmcp")
exec.Command("latte", "install", "libxdmcp")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxfont2")
exec.Command("latte", "install", "libxfont2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: pixman")
exec.Command("latte", "install", "pixman")
	fmt.Println("Instalando dependencia: xauth")
exec.Command("latte", "install", "xauth")
	fmt.Println("Instalando dependencia: xcb-util")
exec.Command("latte", "install", "xcb-util")
	fmt.Println("Instalando dependencia: xcb-util-image")
exec.Command("latte", "install", "xcb-util-image")
	fmt.Println("Instalando dependencia: xcb-util-keysyms")
exec.Command("latte", "install", "xcb-util-keysyms")
	fmt.Println("Instalando dependencia: xcb-util-renderutil")
exec.Command("latte", "install", "xcb-util-renderutil")
	fmt.Println("Instalando dependencia: xcb-util-wm")
exec.Command("latte", "install", "xcb-util-wm")
	fmt.Println("Instalando dependencia: xkbcomp")
exec.Command("latte", "install", "xkbcomp")
	fmt.Println("Instalando dependencia: xkeyboardconfig")
exec.Command("latte", "install", "xkeyboardconfig")
	fmt.Println("Instalando dependencia: libapplewm")
exec.Command("latte", "install", "libapplewm")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: libdrm")
exec.Command("latte", "install", "libdrm")
	fmt.Println("Instalando dependencia: libepoxy")
exec.Command("latte", "install", "libepoxy")
	fmt.Println("Instalando dependencia: libpciaccess")
exec.Command("latte", "install", "libpciaccess")
	fmt.Println("Instalando dependencia: libtirpc")
exec.Command("latte", "install", "libtirpc")
	fmt.Println("Instalando dependencia: libxcvt")
exec.Command("latte", "install", "libxcvt")
	fmt.Println("Instalando dependencia: libxshmfence")
exec.Command("latte", "install", "libxshmfence")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}

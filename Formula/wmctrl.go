package main

// Wmctrl - UNIX/Linux command-line tool to interact with an EWMH/NetWM
// Homepage: https://packages.debian.org/sid/wmctrl

import (
	"fmt"
	
	"os/exec"
)

func installWmctrl() {
	// Método 1: Descargar y extraer .tar.gz
	wmctrl_tar_url := "https://deb.debian.org/debian/pool/main/w/wmctrl/wmctrl_1.07.orig.tar.gz"
	wmctrl_cmd_tar := exec.Command("curl", "-L", wmctrl_tar_url, "-o", "package.tar.gz")
	err := wmctrl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wmctrl_zip_url := "https://deb.debian.org/debian/pool/main/w/wmctrl/wmctrl_1.07.orig.zip"
	wmctrl_cmd_zip := exec.Command("curl", "-L", wmctrl_zip_url, "-o", "package.zip")
	err = wmctrl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wmctrl_bin_url := "https://deb.debian.org/debian/pool/main/w/wmctrl/wmctrl_1.07.orig.bin"
	wmctrl_cmd_bin := exec.Command("curl", "-L", wmctrl_bin_url, "-o", "binary.bin")
	err = wmctrl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wmctrl_src_url := "https://deb.debian.org/debian/pool/main/w/wmctrl/wmctrl_1.07.orig.src.tar.gz"
	wmctrl_cmd_src := exec.Command("curl", "-L", wmctrl_src_url, "-o", "source.tar.gz")
	err = wmctrl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wmctrl_cmd_direct := exec.Command("./binary")
	err = wmctrl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
}

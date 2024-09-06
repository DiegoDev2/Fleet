package main

// SharedMimeInfo - Database of common MIME types
// Homepage: https://wiki.freedesktop.org/www/Software/shared-mime-info

import (
	"fmt"
	
	"os/exec"
)

func installSharedMimeInfo() {
	// Método 1: Descargar y extraer .tar.gz
	sharedmimeinfo_tar_url := "https://gitlab.freedesktop.org/xdg/shared-mime-info/-/archive/2.4/shared-mime-info-2.4.tar.bz2"
	sharedmimeinfo_cmd_tar := exec.Command("curl", "-L", sharedmimeinfo_tar_url, "-o", "package.tar.gz")
	err := sharedmimeinfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sharedmimeinfo_zip_url := "https://gitlab.freedesktop.org/xdg/shared-mime-info/-/archive/2.4/shared-mime-info-2.4.tar.bz2"
	sharedmimeinfo_cmd_zip := exec.Command("curl", "-L", sharedmimeinfo_zip_url, "-o", "package.zip")
	err = sharedmimeinfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sharedmimeinfo_bin_url := "https://gitlab.freedesktop.org/xdg/shared-mime-info/-/archive/2.4/shared-mime-info-2.4.tar.bz2"
	sharedmimeinfo_cmd_bin := exec.Command("curl", "-L", sharedmimeinfo_bin_url, "-o", "binary.bin")
	err = sharedmimeinfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sharedmimeinfo_src_url := "https://gitlab.freedesktop.org/xdg/shared-mime-info/-/archive/2.4/shared-mime-info-2.4.tar.bz2"
	sharedmimeinfo_cmd_src := exec.Command("curl", "-L", sharedmimeinfo_src_url, "-o", "source.tar.gz")
	err = sharedmimeinfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sharedmimeinfo_cmd_direct := exec.Command("./binary")
	err = sharedmimeinfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: itstool")
	exec.Command("latte", "install", "itstool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}

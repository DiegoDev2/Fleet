package main

// Pianod - Pandora client with multiple control interfaces
// Homepage: https://deviousfish.com/pianod/

import (
	"fmt"
	
	"os/exec"
)

func installPianod() {
	// Método 1: Descargar y extraer .tar.gz
	pianod_tar_url := "https://deviousfish.com/Downloads/pianod2/pianod2-405.tar.gz"
	pianod_cmd_tar := exec.Command("curl", "-L", pianod_tar_url, "-o", "package.tar.gz")
	err := pianod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pianod_zip_url := "https://deviousfish.com/Downloads/pianod2/pianod2-405.zip"
	pianod_cmd_zip := exec.Command("curl", "-L", pianod_zip_url, "-o", "package.zip")
	err = pianod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pianod_bin_url := "https://deviousfish.com/Downloads/pianod2/pianod2-405.bin"
	pianod_cmd_bin := exec.Command("curl", "-L", pianod_bin_url, "-o", "binary.bin")
	err = pianod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pianod_src_url := "https://deviousfish.com/Downloads/pianod2/pianod2-405.src.tar.gz"
	pianod_cmd_src := exec.Command("curl", "-L", pianod_src_url, "-o", "source.tar.gz")
	err = pianod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pianod_cmd_direct := exec.Command("./binary")
	err = pianod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gstreamer")
	exec.Command("latte", "install", "gstreamer").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: libao")
	exec.Command("latte", "install", "libao").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: taglib")
	exec.Command("latte", "install", "taglib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libbsd")
	exec.Command("latte", "install", "libbsd").Run()
}

package main

// Procps - Utilities for browsing procfs
// Homepage: https://gitlab.com/procps-ng/procps

import (
	"fmt"
	
	"os/exec"
)

func installProcps() {
	// Método 1: Descargar y extraer .tar.gz
	procps_tar_url := "https://gitlab.com/procps-ng/procps/-/archive/v4.0.4/procps-v4.0.4.tar.gz"
	procps_cmd_tar := exec.Command("curl", "-L", procps_tar_url, "-o", "package.tar.gz")
	err := procps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	procps_zip_url := "https://gitlab.com/procps-ng/procps/-/archive/v4.0.4/procps-v4.0.4.zip"
	procps_cmd_zip := exec.Command("curl", "-L", procps_zip_url, "-o", "package.zip")
	err = procps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	procps_bin_url := "https://gitlab.com/procps-ng/procps/-/archive/v4.0.4/procps-v4.0.4.bin"
	procps_cmd_bin := exec.Command("curl", "-L", procps_bin_url, "-o", "binary.bin")
	err = procps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	procps_src_url := "https://gitlab.com/procps-ng/procps/-/archive/v4.0.4/procps-v4.0.4.src.tar.gz"
	procps_cmd_src := exec.Command("curl", "-L", procps_src_url, "-o", "source.tar.gz")
	err = procps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	procps_cmd_direct := exec.Command("./binary")
	err = procps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
}

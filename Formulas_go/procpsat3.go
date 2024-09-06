package main

// ProcpsAT3 - Utilities for browsing procfs
// Homepage: https://gitlab.com/procps-ng/procps

import (
	"fmt"
	
	"os/exec"
)

func installProcpsAT3() {
	// Método 1: Descargar y extraer .tar.gz
	procpsat3_tar_url := "https://gitlab.com/procps-ng/procps/-/archive/v3.3.17/procps-v3.3.17.tar.gz"
	procpsat3_cmd_tar := exec.Command("curl", "-L", procpsat3_tar_url, "-o", "package.tar.gz")
	err := procpsat3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	procpsat3_zip_url := "https://gitlab.com/procps-ng/procps/-/archive/v3.3.17/procps-v3.3.17.zip"
	procpsat3_cmd_zip := exec.Command("curl", "-L", procpsat3_zip_url, "-o", "package.zip")
	err = procpsat3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	procpsat3_bin_url := "https://gitlab.com/procps-ng/procps/-/archive/v3.3.17/procps-v3.3.17.bin"
	procpsat3_cmd_bin := exec.Command("curl", "-L", procpsat3_bin_url, "-o", "binary.bin")
	err = procpsat3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	procpsat3_src_url := "https://gitlab.com/procps-ng/procps/-/archive/v3.3.17/procps-v3.3.17.src.tar.gz"
	procpsat3_cmd_src := exec.Command("curl", "-L", procpsat3_src_url, "-o", "source.tar.gz")
	err = procpsat3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	procpsat3_cmd_direct := exec.Command("./binary")
	err = procpsat3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}

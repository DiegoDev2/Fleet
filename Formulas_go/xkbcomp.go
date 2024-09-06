package main

// Xkbcomp - XKB keyboard description compiler
// Homepage: https://www.x.org

import (
	"fmt"
	
	"os/exec"
)

func installXkbcomp() {
	// Método 1: Descargar y extraer .tar.gz
	xkbcomp_tar_url := "https://www.x.org/releases/individual/app/xkbcomp-1.4.7.tar.xz"
	xkbcomp_cmd_tar := exec.Command("curl", "-L", xkbcomp_tar_url, "-o", "package.tar.gz")
	err := xkbcomp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xkbcomp_zip_url := "https://www.x.org/releases/individual/app/xkbcomp-1.4.7.tar.xz"
	xkbcomp_cmd_zip := exec.Command("curl", "-L", xkbcomp_zip_url, "-o", "package.zip")
	err = xkbcomp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xkbcomp_bin_url := "https://www.x.org/releases/individual/app/xkbcomp-1.4.7.tar.xz"
	xkbcomp_cmd_bin := exec.Command("curl", "-L", xkbcomp_bin_url, "-o", "binary.bin")
	err = xkbcomp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xkbcomp_src_url := "https://www.x.org/releases/individual/app/xkbcomp-1.4.7.tar.xz"
	xkbcomp_cmd_src := exec.Command("curl", "-L", xkbcomp_src_url, "-o", "source.tar.gz")
	err = xkbcomp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xkbcomp_cmd_direct := exec.Command("./binary")
	err = xkbcomp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxkbfile")
exec.Command("latte", "install", "libxkbfile")
}

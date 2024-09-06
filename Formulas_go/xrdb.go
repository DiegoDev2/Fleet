package main

// Xrdb - X resource database utility
// Homepage: https://gitlab.freedesktop.org/xorg/app/xrdb

import (
	"fmt"
	
	"os/exec"
)

func installXrdb() {
	// Método 1: Descargar y extraer .tar.gz
	xrdb_tar_url := "https://www.x.org/releases/individual/app/xrdb-1.2.2.tar.xz"
	xrdb_cmd_tar := exec.Command("curl", "-L", xrdb_tar_url, "-o", "package.tar.gz")
	err := xrdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xrdb_zip_url := "https://www.x.org/releases/individual/app/xrdb-1.2.2.tar.xz"
	xrdb_cmd_zip := exec.Command("curl", "-L", xrdb_zip_url, "-o", "package.zip")
	err = xrdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xrdb_bin_url := "https://www.x.org/releases/individual/app/xrdb-1.2.2.tar.xz"
	xrdb_cmd_bin := exec.Command("curl", "-L", xrdb_bin_url, "-o", "binary.bin")
	err = xrdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xrdb_src_url := "https://www.x.org/releases/individual/app/xrdb-1.2.2.tar.xz"
	xrdb_cmd_src := exec.Command("curl", "-L", xrdb_src_url, "-o", "source.tar.gz")
	err = xrdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xrdb_cmd_direct := exec.Command("./binary")
	err = xrdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xorg-server")
exec.Command("latte", "install", "xorg-server")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxmu")
exec.Command("latte", "install", "libxmu")
}

package main

// Libagg - High fidelity 2D graphics library for C++
// Homepage: https://antigrain.com/

import (
	"fmt"
	
	"os/exec"
)

func installLibagg() {
	// Método 1: Descargar y extraer .tar.gz
	libagg_tar_url := "https://ftp.osuosl.org/pub/blfs/8.0/a/agg-2.5.tar.gz"
	libagg_cmd_tar := exec.Command("curl", "-L", libagg_tar_url, "-o", "package.tar.gz")
	err := libagg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libagg_zip_url := "https://ftp.osuosl.org/pub/blfs/8.0/a/agg-2.5.zip"
	libagg_cmd_zip := exec.Command("curl", "-L", libagg_zip_url, "-o", "package.zip")
	err = libagg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libagg_bin_url := "https://ftp.osuosl.org/pub/blfs/8.0/a/agg-2.5.bin"
	libagg_cmd_bin := exec.Command("curl", "-L", libagg_bin_url, "-o", "binary.bin")
	err = libagg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libagg_src_url := "https://ftp.osuosl.org/pub/blfs/8.0/a/agg-2.5.src.tar.gz"
	libagg_cmd_src := exec.Command("curl", "-L", libagg_src_url, "-o", "source.tar.gz")
	err = libagg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libagg_cmd_direct := exec.Command("./binary")
	err = libagg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
}

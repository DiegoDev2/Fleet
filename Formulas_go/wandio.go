package main

// Wandio - Transparently read from and write to zip, bzip2, lzma or zstd archives
// Homepage: https://github.com/LibtraceTeam/wandio

import (
	"fmt"
	
	"os/exec"
)

func installWandio() {
	// Método 1: Descargar y extraer .tar.gz
	wandio_tar_url := "https://github.com/LibtraceTeam/wandio/archive/refs/tags/4.2.6-1.tar.gz"
	wandio_cmd_tar := exec.Command("curl", "-L", wandio_tar_url, "-o", "package.tar.gz")
	err := wandio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wandio_zip_url := "https://github.com/LibtraceTeam/wandio/archive/refs/tags/4.2.6-1.zip"
	wandio_cmd_zip := exec.Command("curl", "-L", wandio_zip_url, "-o", "package.zip")
	err = wandio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wandio_bin_url := "https://github.com/LibtraceTeam/wandio/archive/refs/tags/4.2.6-1.bin"
	wandio_cmd_bin := exec.Command("curl", "-L", wandio_bin_url, "-o", "binary.bin")
	err = wandio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wandio_src_url := "https://github.com/LibtraceTeam/wandio/archive/refs/tags/4.2.6-1.src.tar.gz"
	wandio_cmd_src := exec.Command("curl", "-L", wandio_src_url, "-o", "source.tar.gz")
	err = wandio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wandio_cmd_direct := exec.Command("./binary")
	err = wandio_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}

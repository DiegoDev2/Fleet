package main

// Webfs - HTTP server for purely static content
// Homepage: https://linux.bytesex.org/misc/webfs.html

import (
	"fmt"
	
	"os/exec"
)

func installWebfs() {
	// Método 1: Descargar y extraer .tar.gz
	webfs_tar_url := "https://www.kraxel.org/releases/webfs/webfs-1.21.tar.gz"
	webfs_cmd_tar := exec.Command("curl", "-L", webfs_tar_url, "-o", "package.tar.gz")
	err := webfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webfs_zip_url := "https://www.kraxel.org/releases/webfs/webfs-1.21.zip"
	webfs_cmd_zip := exec.Command("curl", "-L", webfs_zip_url, "-o", "package.zip")
	err = webfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webfs_bin_url := "https://www.kraxel.org/releases/webfs/webfs-1.21.bin"
	webfs_cmd_bin := exec.Command("curl", "-L", webfs_bin_url, "-o", "binary.bin")
	err = webfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webfs_src_url := "https://www.kraxel.org/releases/webfs/webfs-1.21.src.tar.gz"
	webfs_cmd_src := exec.Command("curl", "-L", webfs_src_url, "-o", "source.tar.gz")
	err = webfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webfs_cmd_direct := exec.Command("./binary")
	err = webfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: httpd")
exec.Command("latte", "install", "httpd")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

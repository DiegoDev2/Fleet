package main

// Gocryptfs - Encrypted overlay filesystem written in Go
// Homepage: https://nuetzlich.net/gocryptfs/

import (
	"fmt"
	
	"os/exec"
)

func installGocryptfs() {
	// Método 1: Descargar y extraer .tar.gz
	gocryptfs_tar_url := "https://github.com/rfjakob/gocryptfs/releases/download/v2.4.0/gocryptfs_v2.4.0_src-deps.tar.gz"
	gocryptfs_cmd_tar := exec.Command("curl", "-L", gocryptfs_tar_url, "-o", "package.tar.gz")
	err := gocryptfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gocryptfs_zip_url := "https://github.com/rfjakob/gocryptfs/releases/download/v2.4.0/gocryptfs_v2.4.0_src-deps.zip"
	gocryptfs_cmd_zip := exec.Command("curl", "-L", gocryptfs_zip_url, "-o", "package.zip")
	err = gocryptfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gocryptfs_bin_url := "https://github.com/rfjakob/gocryptfs/releases/download/v2.4.0/gocryptfs_v2.4.0_src-deps.bin"
	gocryptfs_cmd_bin := exec.Command("curl", "-L", gocryptfs_bin_url, "-o", "binary.bin")
	err = gocryptfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gocryptfs_src_url := "https://github.com/rfjakob/gocryptfs/releases/download/v2.4.0/gocryptfs_v2.4.0_src-deps.src.tar.gz"
	gocryptfs_cmd_src := exec.Command("curl", "-L", gocryptfs_src_url, "-o", "source.tar.gz")
	err = gocryptfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gocryptfs_cmd_direct := exec.Command("./binary")
	err = gocryptfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libfuse")
exec.Command("latte", "install", "libfuse")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

package main

// Felinks - Text mode browser and Gemini, NNTP, FTP, Gopher, Finger, and BitTorrent client
// Homepage: https://github.com/rkd77/elinks

import (
	"fmt"
	
	"os/exec"
)

func installFelinks() {
	// Método 1: Descargar y extraer .tar.gz
	felinks_tar_url := "https://github.com/rkd77/elinks/releases/download/v0.17.0/elinks-0.17.0.tar.xz"
	felinks_cmd_tar := exec.Command("curl", "-L", felinks_tar_url, "-o", "package.tar.gz")
	err := felinks_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	felinks_zip_url := "https://github.com/rkd77/elinks/releases/download/v0.17.0/elinks-0.17.0.tar.xz"
	felinks_cmd_zip := exec.Command("curl", "-L", felinks_zip_url, "-o", "package.zip")
	err = felinks_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	felinks_bin_url := "https://github.com/rkd77/elinks/releases/download/v0.17.0/elinks-0.17.0.tar.xz"
	felinks_cmd_bin := exec.Command("curl", "-L", felinks_bin_url, "-o", "binary.bin")
	err = felinks_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	felinks_src_url := "https://github.com/rkd77/elinks/releases/download/v0.17.0/elinks-0.17.0.tar.xz"
	felinks_cmd_src := exec.Command("curl", "-L", felinks_src_url, "-o", "source.tar.gz")
	err = felinks_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	felinks_cmd_direct := exec.Command("./binary")
	err = felinks_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: brotli")
	exec.Command("latte", "install", "brotli").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: tre")
	exec.Command("latte", "install", "tre").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}

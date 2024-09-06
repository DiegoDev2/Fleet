package main

// Sh4d0wup - Signing-key abuse and update exploitation framework
// Homepage: https://github.com/kpcyrd/sh4d0wup

import (
	"fmt"
	
	"os/exec"
)

func installSh4d0wup() {
	// Método 1: Descargar y extraer .tar.gz
	sh4d0wup_tar_url := "https://github.com/kpcyrd/sh4d0wup/archive/refs/tags/v0.9.3.tar.gz"
	sh4d0wup_cmd_tar := exec.Command("curl", "-L", sh4d0wup_tar_url, "-o", "package.tar.gz")
	err := sh4d0wup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sh4d0wup_zip_url := "https://github.com/kpcyrd/sh4d0wup/archive/refs/tags/v0.9.3.zip"
	sh4d0wup_cmd_zip := exec.Command("curl", "-L", sh4d0wup_zip_url, "-o", "package.zip")
	err = sh4d0wup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sh4d0wup_bin_url := "https://github.com/kpcyrd/sh4d0wup/archive/refs/tags/v0.9.3.bin"
	sh4d0wup_cmd_bin := exec.Command("curl", "-L", sh4d0wup_bin_url, "-o", "binary.bin")
	err = sh4d0wup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sh4d0wup_src_url := "https://github.com/kpcyrd/sh4d0wup/archive/refs/tags/v0.9.3.src.tar.gz"
	sh4d0wup_cmd_src := exec.Command("curl", "-L", sh4d0wup_src_url, "-o", "source.tar.gz")
	err = sh4d0wup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sh4d0wup_cmd_direct := exec.Command("./binary")
	err = sh4d0wup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pgpdump")
	exec.Command("latte", "install", "pgpdump").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcsc-lite")
	exec.Command("latte", "install", "pcsc-lite").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}

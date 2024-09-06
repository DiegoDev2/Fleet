package main

// Moarvm - VM with adaptive optimization and JIT compilation, built for Rakudo
// Homepage: https://moarvm.org

import (
	"fmt"
	
	"os/exec"
)

func installMoarvm() {
	// Método 1: Descargar y extraer .tar.gz
	moarvm_tar_url := "https://github.com/MoarVM/MoarVM/releases/download/2024.08/MoarVM-2024.08.tar.gz"
	moarvm_cmd_tar := exec.Command("curl", "-L", moarvm_tar_url, "-o", "package.tar.gz")
	err := moarvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	moarvm_zip_url := "https://github.com/MoarVM/MoarVM/releases/download/2024.08/MoarVM-2024.08.zip"
	moarvm_cmd_zip := exec.Command("curl", "-L", moarvm_zip_url, "-o", "package.zip")
	err = moarvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	moarvm_bin_url := "https://github.com/MoarVM/MoarVM/releases/download/2024.08/MoarVM-2024.08.bin"
	moarvm_cmd_bin := exec.Command("curl", "-L", moarvm_bin_url, "-o", "binary.bin")
	err = moarvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	moarvm_src_url := "https://github.com/MoarVM/MoarVM/releases/download/2024.08/MoarVM-2024.08.src.tar.gz"
	moarvm_cmd_src := exec.Command("curl", "-L", moarvm_src_url, "-o", "source.tar.gz")
	err = moarvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	moarvm_cmd_direct := exec.Command("./binary")
	err = moarvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtommath")
	exec.Command("latte", "install", "libtommath").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}

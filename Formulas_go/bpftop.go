package main

// Bpftop - Dynamic real-time view of running eBPF programs
// Homepage: https://github.com/Netflix/bpftop

import (
	"fmt"
	
	"os/exec"
)

func installBpftop() {
	// Método 1: Descargar y extraer .tar.gz
	bpftop_tar_url := "https://github.com/Netflix/bpftop/archive/refs/tags/v0.5.2.tar.gz"
	bpftop_cmd_tar := exec.Command("curl", "-L", bpftop_tar_url, "-o", "package.tar.gz")
	err := bpftop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bpftop_zip_url := "https://github.com/Netflix/bpftop/archive/refs/tags/v0.5.2.zip"
	bpftop_cmd_zip := exec.Command("curl", "-L", bpftop_zip_url, "-o", "package.zip")
	err = bpftop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bpftop_bin_url := "https://github.com/Netflix/bpftop/archive/refs/tags/v0.5.2.bin"
	bpftop_cmd_bin := exec.Command("curl", "-L", bpftop_bin_url, "-o", "binary.bin")
	err = bpftop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bpftop_src_url := "https://github.com/Netflix/bpftop/archive/refs/tags/v0.5.2.src.tar.gz"
	bpftop_cmd_src := exec.Command("curl", "-L", bpftop_src_url, "-o", "source.tar.gz")
	err = bpftop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bpftop_cmd_direct := exec.Command("./binary")
	err = bpftop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
}

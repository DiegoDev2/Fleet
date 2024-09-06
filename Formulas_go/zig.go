package main

// Zig - Programming language designed for robustness, optimality, and clarity
// Homepage: https://ziglang.org/

import (
	"fmt"
	
	"os/exec"
)

func installZig() {
	// Método 1: Descargar y extraer .tar.gz
	zig_tar_url := "https://ziglang.org/download/0.13.0/zig-0.13.0.tar.xz"
	zig_cmd_tar := exec.Command("curl", "-L", zig_tar_url, "-o", "package.tar.gz")
	err := zig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zig_zip_url := "https://ziglang.org/download/0.13.0/zig-0.13.0.tar.xz"
	zig_cmd_zip := exec.Command("curl", "-L", zig_zip_url, "-o", "package.zip")
	err = zig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zig_bin_url := "https://ziglang.org/download/0.13.0/zig-0.13.0.tar.xz"
	zig_cmd_bin := exec.Command("curl", "-L", zig_bin_url, "-o", "binary.bin")
	err = zig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zig_src_url := "https://ziglang.org/download/0.13.0/zig-0.13.0.tar.xz"
	zig_cmd_src := exec.Command("curl", "-L", zig_src_url, "-o", "source.tar.gz")
	err = zig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zig_cmd_direct := exec.Command("./binary")
	err = zig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: z3")
exec.Command("latte", "install", "z3")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}

package main

// GccAT5 - GNU Compiler Collection
// Homepage: https://gcc.gnu.org/

import (
	"fmt"
	
	"os/exec"
)

func installGccAT5() {
	// Método 1: Descargar y extraer .tar.gz
	gccat5_tar_url := "https://ftp.gnu.org/gnu/gcc/gcc-5.5.0/gcc-5.5.0.tar.xz"
	gccat5_cmd_tar := exec.Command("curl", "-L", gccat5_tar_url, "-o", "package.tar.gz")
	err := gccat5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gccat5_zip_url := "https://ftp.gnu.org/gnu/gcc/gcc-5.5.0/gcc-5.5.0.tar.xz"
	gccat5_cmd_zip := exec.Command("curl", "-L", gccat5_zip_url, "-o", "package.zip")
	err = gccat5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gccat5_bin_url := "https://ftp.gnu.org/gnu/gcc/gcc-5.5.0/gcc-5.5.0.tar.xz"
	gccat5_cmd_bin := exec.Command("curl", "-L", gccat5_bin_url, "-o", "binary.bin")
	err = gccat5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gccat5_src_url := "https://ftp.gnu.org/gnu/gcc/gcc-5.5.0/gcc-5.5.0.tar.xz"
	gccat5_cmd_src := exec.Command("curl", "-L", gccat5_src_url, "-o", "source.tar.gz")
	err = gccat5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gccat5_cmd_direct := exec.Command("./binary")
	err = gccat5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: isl@0.18")
exec.Command("latte", "install", "isl@0.18")
	fmt.Println("Instalando dependencia: libmpc")
exec.Command("latte", "install", "libmpc")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: binutils")
exec.Command("latte", "install", "binutils")
}

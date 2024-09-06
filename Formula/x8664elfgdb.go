package main

// X8664ElfGdb - GNU debugger for x86_64-elf cross development
// Homepage: https://www.gnu.org/software/gdb/

import (
	"fmt"
	
	"os/exec"
)

func installX8664ElfGdb() {
	// Método 1: Descargar y extraer .tar.gz
	x8664elfgdb_tar_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	x8664elfgdb_cmd_tar := exec.Command("curl", "-L", x8664elfgdb_tar_url, "-o", "package.tar.gz")
	err := x8664elfgdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	x8664elfgdb_zip_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	x8664elfgdb_cmd_zip := exec.Command("curl", "-L", x8664elfgdb_zip_url, "-o", "package.zip")
	err = x8664elfgdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	x8664elfgdb_bin_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	x8664elfgdb_cmd_bin := exec.Command("curl", "-L", x8664elfgdb_bin_url, "-o", "binary.bin")
	err = x8664elfgdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	x8664elfgdb_src_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	x8664elfgdb_cmd_src := exec.Command("curl", "-L", x8664elfgdb_src_url, "-o", "source.tar.gz")
	err = x8664elfgdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	x8664elfgdb_cmd_direct := exec.Command("./binary")
	err = x8664elfgdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: x86_64-elf-gcc")
	exec.Command("latte", "install", "x86_64-elf-gcc").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}

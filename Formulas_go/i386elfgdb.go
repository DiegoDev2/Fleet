package main

// I386ElfGdb - GNU debugger for i386-elf cross development
// Homepage: https://www.gnu.org/software/gdb/

import (
	"fmt"
	
	"os/exec"
)

func installI386ElfGdb() {
	// Método 1: Descargar y extraer .tar.gz
	i386elfgdb_tar_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	i386elfgdb_cmd_tar := exec.Command("curl", "-L", i386elfgdb_tar_url, "-o", "package.tar.gz")
	err := i386elfgdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	i386elfgdb_zip_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	i386elfgdb_cmd_zip := exec.Command("curl", "-L", i386elfgdb_zip_url, "-o", "package.zip")
	err = i386elfgdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	i386elfgdb_bin_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	i386elfgdb_cmd_bin := exec.Command("curl", "-L", i386elfgdb_bin_url, "-o", "binary.bin")
	err = i386elfgdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	i386elfgdb_src_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	i386elfgdb_cmd_src := exec.Command("curl", "-L", i386elfgdb_src_url, "-o", "source.tar.gz")
	err = i386elfgdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	i386elfgdb_cmd_direct := exec.Command("./binary")
	err = i386elfgdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: i686-elf-gcc")
exec.Command("latte", "install", "i686-elf-gcc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}

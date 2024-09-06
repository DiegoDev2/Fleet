package main

// ArmNoneEabiGdb - GNU debugger for arm-none-eabi cross development
// Homepage: https://www.gnu.org/software/gdb/

import (
	"fmt"
	
	"os/exec"
)

func installArmNoneEabiGdb() {
	// Método 1: Descargar y extraer .tar.gz
	armnoneeabigdb_tar_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	armnoneeabigdb_cmd_tar := exec.Command("curl", "-L", armnoneeabigdb_tar_url, "-o", "package.tar.gz")
	err := armnoneeabigdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	armnoneeabigdb_zip_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	armnoneeabigdb_cmd_zip := exec.Command("curl", "-L", armnoneeabigdb_zip_url, "-o", "package.zip")
	err = armnoneeabigdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	armnoneeabigdb_bin_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	armnoneeabigdb_cmd_bin := exec.Command("curl", "-L", armnoneeabigdb_bin_url, "-o", "binary.bin")
	err = armnoneeabigdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	armnoneeabigdb_src_url := "https://ftp.gnu.org/gnu/gdb/gdb-15.1.tar.xz"
	armnoneeabigdb_cmd_src := exec.Command("curl", "-L", armnoneeabigdb_src_url, "-o", "source.tar.gz")
	err = armnoneeabigdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	armnoneeabigdb_cmd_direct := exec.Command("./binary")
	err = armnoneeabigdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: arm-none-eabi-gcc")
	exec.Command("latte", "install", "arm-none-eabi-gcc").Run()
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

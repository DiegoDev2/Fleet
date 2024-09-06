package main

// ArmNoneEabiGcc - GNU compiler collection for arm-none-eabi
// Homepage: https://gcc.gnu.org

import (
	"fmt"
	
	"os/exec"
)

func installArmNoneEabiGcc() {
	// Método 1: Descargar y extraer .tar.gz
	armnoneeabigcc_tar_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	armnoneeabigcc_cmd_tar := exec.Command("curl", "-L", armnoneeabigcc_tar_url, "-o", "package.tar.gz")
	err := armnoneeabigcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	armnoneeabigcc_zip_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	armnoneeabigcc_cmd_zip := exec.Command("curl", "-L", armnoneeabigcc_zip_url, "-o", "package.zip")
	err = armnoneeabigcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	armnoneeabigcc_bin_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	armnoneeabigcc_cmd_bin := exec.Command("curl", "-L", armnoneeabigcc_bin_url, "-o", "binary.bin")
	err = armnoneeabigcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	armnoneeabigcc_src_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	armnoneeabigcc_cmd_src := exec.Command("curl", "-L", armnoneeabigcc_src_url, "-o", "source.tar.gz")
	err = armnoneeabigcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	armnoneeabigcc_cmd_direct := exec.Command("./binary")
	err = armnoneeabigcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: arm-none-eabi-binutils")
	exec.Command("latte", "install", "arm-none-eabi-binutils").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: isl")
	exec.Command("latte", "install", "isl").Run()
	fmt.Println("Instalando dependencia: libmpc")
	exec.Command("latte", "install", "libmpc").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}

package main

// ArmLinuxGnueabihfBinutils - FSF/GNU binutils for cross-compiling to arm-linux
// Homepage: https://www.gnu.org/software/binutils/binutils.html

import (
	"fmt"
	
	"os/exec"
)

func installArmLinuxGnueabihfBinutils() {
	// Método 1: Descargar y extraer .tar.gz
	armlinuxgnueabihfbinutils_tar_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	armlinuxgnueabihfbinutils_cmd_tar := exec.Command("curl", "-L", armlinuxgnueabihfbinutils_tar_url, "-o", "package.tar.gz")
	err := armlinuxgnueabihfbinutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	armlinuxgnueabihfbinutils_zip_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	armlinuxgnueabihfbinutils_cmd_zip := exec.Command("curl", "-L", armlinuxgnueabihfbinutils_zip_url, "-o", "package.zip")
	err = armlinuxgnueabihfbinutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	armlinuxgnueabihfbinutils_bin_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	armlinuxgnueabihfbinutils_cmd_bin := exec.Command("curl", "-L", armlinuxgnueabihfbinutils_bin_url, "-o", "binary.bin")
	err = armlinuxgnueabihfbinutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	armlinuxgnueabihfbinutils_src_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	armlinuxgnueabihfbinutils_cmd_src := exec.Command("curl", "-L", armlinuxgnueabihfbinutils_src_url, "-o", "source.tar.gz")
	err = armlinuxgnueabihfbinutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	armlinuxgnueabihfbinutils_cmd_direct := exec.Command("./binary")
	err = armlinuxgnueabihfbinutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}

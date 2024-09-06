package main

// Mpfi - Multiple precision interval arithmetic library
// Homepage: https://perso.ens-lyon.fr/nathalie.revol/software.html

import (
	"fmt"
	
	"os/exec"
)

func installMpfi() {
	// Método 1: Descargar y extraer .tar.gz
	mpfi_tar_url := "https://perso.ens-lyon.fr/nathalie.revol/softwares/mpfi-1.5.4.tar.bz2"
	mpfi_cmd_tar := exec.Command("curl", "-L", mpfi_tar_url, "-o", "package.tar.gz")
	err := mpfi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpfi_zip_url := "https://perso.ens-lyon.fr/nathalie.revol/softwares/mpfi-1.5.4.tar.bz2"
	mpfi_cmd_zip := exec.Command("curl", "-L", mpfi_zip_url, "-o", "package.zip")
	err = mpfi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpfi_bin_url := "https://perso.ens-lyon.fr/nathalie.revol/softwares/mpfi-1.5.4.tar.bz2"
	mpfi_cmd_bin := exec.Command("curl", "-L", mpfi_bin_url, "-o", "binary.bin")
	err = mpfi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpfi_src_url := "https://perso.ens-lyon.fr/nathalie.revol/softwares/mpfi-1.5.4.tar.bz2"
	mpfi_cmd_src := exec.Command("curl", "-L", mpfi_src_url, "-o", "source.tar.gz")
	err = mpfi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpfi_cmd_direct := exec.Command("./binary")
	err = mpfi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
}

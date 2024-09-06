package main

// IsoCodes - Provides lists of various ISO standards
// Homepage: https://salsa.debian.org/iso-codes-team/iso-codes

import (
	"fmt"
	
	"os/exec"
)

func installIsoCodes() {
	// Método 1: Descargar y extraer .tar.gz
	isocodes_tar_url := "https://deb.debian.org/debian/pool/main/i/iso-codes/iso-codes_4.16.0.orig.tar.xz"
	isocodes_cmd_tar := exec.Command("curl", "-L", isocodes_tar_url, "-o", "package.tar.gz")
	err := isocodes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	isocodes_zip_url := "https://deb.debian.org/debian/pool/main/i/iso-codes/iso-codes_4.16.0.orig.tar.xz"
	isocodes_cmd_zip := exec.Command("curl", "-L", isocodes_zip_url, "-o", "package.zip")
	err = isocodes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	isocodes_bin_url := "https://deb.debian.org/debian/pool/main/i/iso-codes/iso-codes_4.16.0.orig.tar.xz"
	isocodes_cmd_bin := exec.Command("curl", "-L", isocodes_bin_url, "-o", "binary.bin")
	err = isocodes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	isocodes_src_url := "https://deb.debian.org/debian/pool/main/i/iso-codes/iso-codes_4.16.0.orig.tar.xz"
	isocodes_cmd_src := exec.Command("curl", "-L", isocodes_src_url, "-o", "source.tar.gz")
	err = isocodes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	isocodes_cmd_direct := exec.Command("./binary")
	err = isocodes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}

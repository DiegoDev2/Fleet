package main

// Ats2Postiats - Programming language with formal specification features
// Homepage: http://www.ats-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installAts2Postiats() {
	// Método 1: Descargar y extraer .tar.gz
	ats2postiats_tar_url := "https://downloads.sourceforge.net/project/ats2-lang/ats2-lang/ats2-postiats-0.4.2/ATS2-Postiats-0.4.2.tgz"
	ats2postiats_cmd_tar := exec.Command("curl", "-L", ats2postiats_tar_url, "-o", "package.tar.gz")
	err := ats2postiats_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ats2postiats_zip_url := "https://downloads.sourceforge.net/project/ats2-lang/ats2-lang/ats2-postiats-0.4.2/ATS2-Postiats-0.4.2.tgz"
	ats2postiats_cmd_zip := exec.Command("curl", "-L", ats2postiats_zip_url, "-o", "package.zip")
	err = ats2postiats_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ats2postiats_bin_url := "https://downloads.sourceforge.net/project/ats2-lang/ats2-lang/ats2-postiats-0.4.2/ATS2-Postiats-0.4.2.tgz"
	ats2postiats_cmd_bin := exec.Command("curl", "-L", ats2postiats_bin_url, "-o", "binary.bin")
	err = ats2postiats_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ats2postiats_src_url := "https://downloads.sourceforge.net/project/ats2-lang/ats2-lang/ats2-postiats-0.4.2/ATS2-Postiats-0.4.2.tgz"
	ats2postiats_cmd_src := exec.Command("curl", "-L", ats2postiats_src_url, "-o", "source.tar.gz")
	err = ats2postiats_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ats2postiats_cmd_direct := exec.Command("./binary")
	err = ats2postiats_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}

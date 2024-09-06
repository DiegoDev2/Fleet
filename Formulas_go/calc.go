package main

// Calc - Arbitrary precision calculator
// Homepage: http://www.isthe.com/chongo/tech/comp/calc/

import (
	"fmt"
	
	"os/exec"
)

func installCalc() {
	// Método 1: Descargar y extraer .tar.gz
	calc_tar_url := "https://downloads.sourceforge.net/project/calc/calc/2.15.1.0/calc-2.15.1.0.tar.bz2"
	calc_cmd_tar := exec.Command("curl", "-L", calc_tar_url, "-o", "package.tar.gz")
	err := calc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	calc_zip_url := "https://downloads.sourceforge.net/project/calc/calc/2.15.1.0/calc-2.15.1.0.tar.bz2"
	calc_cmd_zip := exec.Command("curl", "-L", calc_zip_url, "-o", "package.zip")
	err = calc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	calc_bin_url := "https://downloads.sourceforge.net/project/calc/calc/2.15.1.0/calc-2.15.1.0.tar.bz2"
	calc_cmd_bin := exec.Command("curl", "-L", calc_bin_url, "-o", "binary.bin")
	err = calc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	calc_src_url := "https://downloads.sourceforge.net/project/calc/calc/2.15.1.0/calc-2.15.1.0.tar.bz2"
	calc_cmd_src := exec.Command("curl", "-L", calc_src_url, "-o", "source.tar.gz")
	err = calc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	calc_cmd_direct := exec.Command("./binary")
	err = calc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}

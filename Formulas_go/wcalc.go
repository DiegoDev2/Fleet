package main

// WCalc - Very capable calculator
// Homepage: https://w-calc.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installWCalc() {
	// Método 1: Descargar y extraer .tar.gz
	wcalc_tar_url := "https://downloads.sourceforge.net/project/w-calc/Wcalc/2.5/wcalc-2.5.tar.bz2"
	wcalc_cmd_tar := exec.Command("curl", "-L", wcalc_tar_url, "-o", "package.tar.gz")
	err := wcalc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wcalc_zip_url := "https://downloads.sourceforge.net/project/w-calc/Wcalc/2.5/wcalc-2.5.tar.bz2"
	wcalc_cmd_zip := exec.Command("curl", "-L", wcalc_zip_url, "-o", "package.zip")
	err = wcalc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wcalc_bin_url := "https://downloads.sourceforge.net/project/w-calc/Wcalc/2.5/wcalc-2.5.tar.bz2"
	wcalc_cmd_bin := exec.Command("curl", "-L", wcalc_bin_url, "-o", "binary.bin")
	err = wcalc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wcalc_src_url := "https://downloads.sourceforge.net/project/w-calc/Wcalc/2.5/wcalc-2.5.tar.bz2"
	wcalc_cmd_src := exec.Command("curl", "-L", wcalc_src_url, "-o", "source.tar.gz")
	err = wcalc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wcalc_cmd_direct := exec.Command("./binary")
	err = wcalc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}

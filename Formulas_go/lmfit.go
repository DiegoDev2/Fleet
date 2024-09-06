package main

// Lmfit - C library for Levenberg-Marquardt minimization and least-squares fitting
// Homepage: https://jugit.fz-juelich.de/mlz/lmfit

import (
	"fmt"
	
	"os/exec"
)

func installLmfit() {
	// Método 1: Descargar y extraer .tar.gz
	lmfit_tar_url := "https://jugit.fz-juelich.de/mlz/lmfit/-/archive/v9.0/lmfit-v9.0.tar.bz2"
	lmfit_cmd_tar := exec.Command("curl", "-L", lmfit_tar_url, "-o", "package.tar.gz")
	err := lmfit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lmfit_zip_url := "https://jugit.fz-juelich.de/mlz/lmfit/-/archive/v9.0/lmfit-v9.0.tar.bz2"
	lmfit_cmd_zip := exec.Command("curl", "-L", lmfit_zip_url, "-o", "package.zip")
	err = lmfit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lmfit_bin_url := "https://jugit.fz-juelich.de/mlz/lmfit/-/archive/v9.0/lmfit-v9.0.tar.bz2"
	lmfit_cmd_bin := exec.Command("curl", "-L", lmfit_bin_url, "-o", "binary.bin")
	err = lmfit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lmfit_src_url := "https://jugit.fz-juelich.de/mlz/lmfit/-/archive/v9.0/lmfit-v9.0.tar.bz2"
	lmfit_cmd_src := exec.Command("curl", "-L", lmfit_src_url, "-o", "source.tar.gz")
	err = lmfit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lmfit_cmd_direct := exec.Command("./binary")
	err = lmfit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

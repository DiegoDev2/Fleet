package main

// DatetimeFortran - Fortran time and date manipulation library
// Homepage: https://github.com/wavebitscientific/datetime-fortran

import (
	"fmt"
	
	"os/exec"
)

func installDatetimeFortran() {
	// Método 1: Descargar y extraer .tar.gz
	datetimefortran_tar_url := "https://github.com/wavebitscientific/datetime-fortran/releases/download/v1.7.0/datetime-fortran-1.7.0.tar.gz"
	datetimefortran_cmd_tar := exec.Command("curl", "-L", datetimefortran_tar_url, "-o", "package.tar.gz")
	err := datetimefortran_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	datetimefortran_zip_url := "https://github.com/wavebitscientific/datetime-fortran/releases/download/v1.7.0/datetime-fortran-1.7.0.zip"
	datetimefortran_cmd_zip := exec.Command("curl", "-L", datetimefortran_zip_url, "-o", "package.zip")
	err = datetimefortran_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	datetimefortran_bin_url := "https://github.com/wavebitscientific/datetime-fortran/releases/download/v1.7.0/datetime-fortran-1.7.0.bin"
	datetimefortran_cmd_bin := exec.Command("curl", "-L", datetimefortran_bin_url, "-o", "binary.bin")
	err = datetimefortran_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	datetimefortran_src_url := "https://github.com/wavebitscientific/datetime-fortran/releases/download/v1.7.0/datetime-fortran-1.7.0.src.tar.gz"
	datetimefortran_cmd_src := exec.Command("curl", "-L", datetimefortran_src_url, "-o", "source.tar.gz")
	err = datetimefortran_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	datetimefortran_cmd_direct := exec.Command("./binary")
	err = datetimefortran_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}

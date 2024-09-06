package main

// HowardHinnantDate - C++ library for date and time operations based on <chrono>
// Homepage: https://github.com/HowardHinnant/date

import (
	"fmt"
	
	"os/exec"
)

func installHowardHinnantDate() {
	// Método 1: Descargar y extraer .tar.gz
	howardhinnantdate_tar_url := "https://github.com/HowardHinnant/date/archive/refs/tags/v3.0.1.tar.gz"
	howardhinnantdate_cmd_tar := exec.Command("curl", "-L", howardhinnantdate_tar_url, "-o", "package.tar.gz")
	err := howardhinnantdate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	howardhinnantdate_zip_url := "https://github.com/HowardHinnant/date/archive/refs/tags/v3.0.1.zip"
	howardhinnantdate_cmd_zip := exec.Command("curl", "-L", howardhinnantdate_zip_url, "-o", "package.zip")
	err = howardhinnantdate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	howardhinnantdate_bin_url := "https://github.com/HowardHinnant/date/archive/refs/tags/v3.0.1.bin"
	howardhinnantdate_cmd_bin := exec.Command("curl", "-L", howardhinnantdate_bin_url, "-o", "binary.bin")
	err = howardhinnantdate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	howardhinnantdate_src_url := "https://github.com/HowardHinnant/date/archive/refs/tags/v3.0.1.src.tar.gz"
	howardhinnantdate_cmd_src := exec.Command("curl", "-L", howardhinnantdate_src_url, "-o", "source.tar.gz")
	err = howardhinnantdate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	howardhinnantdate_cmd_direct := exec.Command("./binary")
	err = howardhinnantdate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}

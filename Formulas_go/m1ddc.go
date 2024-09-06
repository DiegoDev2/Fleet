package main

// M1ddc - Control external displays (USB-C/DisplayPort Alt Mode) using DDC/CI on M1 Macs
// Homepage: https://github.com/waydabber/m1ddc

import (
	"fmt"
	
	"os/exec"
)

func installM1ddc() {
	// Método 1: Descargar y extraer .tar.gz
	m1ddc_tar_url := "https://github.com/waydabber/m1ddc/archive/refs/tags/v1.2.0.tar.gz"
	m1ddc_cmd_tar := exec.Command("curl", "-L", m1ddc_tar_url, "-o", "package.tar.gz")
	err := m1ddc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	m1ddc_zip_url := "https://github.com/waydabber/m1ddc/archive/refs/tags/v1.2.0.zip"
	m1ddc_cmd_zip := exec.Command("curl", "-L", m1ddc_zip_url, "-o", "package.zip")
	err = m1ddc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	m1ddc_bin_url := "https://github.com/waydabber/m1ddc/archive/refs/tags/v1.2.0.bin"
	m1ddc_cmd_bin := exec.Command("curl", "-L", m1ddc_bin_url, "-o", "binary.bin")
	err = m1ddc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	m1ddc_src_url := "https://github.com/waydabber/m1ddc/archive/refs/tags/v1.2.0.src.tar.gz"
	m1ddc_cmd_src := exec.Command("curl", "-L", m1ddc_src_url, "-o", "source.tar.gz")
	err = m1ddc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	m1ddc_cmd_direct := exec.Command("./binary")
	err = m1ddc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}

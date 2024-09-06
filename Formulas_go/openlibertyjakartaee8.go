package main

// OpenlibertyJakartaee8 - Lightweight open framework for Java (Jakarta EE 8)
// Homepage: https://openliberty.io

import (
	"fmt"
	
	"os/exec"
)

func installOpenlibertyJakartaee8() {
	// Método 1: Descargar y extraer .tar.gz
	openlibertyjakartaee8_tar_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-javaee8-24.0.0.8.zip"
	openlibertyjakartaee8_cmd_tar := exec.Command("curl", "-L", openlibertyjakartaee8_tar_url, "-o", "package.tar.gz")
	err := openlibertyjakartaee8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openlibertyjakartaee8_zip_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-javaee8-24.0.0.8.zip"
	openlibertyjakartaee8_cmd_zip := exec.Command("curl", "-L", openlibertyjakartaee8_zip_url, "-o", "package.zip")
	err = openlibertyjakartaee8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openlibertyjakartaee8_bin_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-javaee8-24.0.0.8.zip"
	openlibertyjakartaee8_cmd_bin := exec.Command("curl", "-L", openlibertyjakartaee8_bin_url, "-o", "binary.bin")
	err = openlibertyjakartaee8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openlibertyjakartaee8_src_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-javaee8-24.0.0.8.zip"
	openlibertyjakartaee8_cmd_src := exec.Command("curl", "-L", openlibertyjakartaee8_src_url, "-o", "source.tar.gz")
	err = openlibertyjakartaee8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openlibertyjakartaee8_cmd_direct := exec.Command("./binary")
	err = openlibertyjakartaee8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}

package main

// OpenlibertyWebprofile8 - Lightweight open framework for Java (Jakarta EE Web Profile 8)
// Homepage: https://openliberty.io

import (
	"fmt"
	
	"os/exec"
)

func installOpenlibertyWebprofile8() {
	// Método 1: Descargar y extraer .tar.gz
	openlibertywebprofile8_tar_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-webProfile8-24.0.0.8.zip"
	openlibertywebprofile8_cmd_tar := exec.Command("curl", "-L", openlibertywebprofile8_tar_url, "-o", "package.tar.gz")
	err := openlibertywebprofile8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openlibertywebprofile8_zip_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-webProfile8-24.0.0.8.zip"
	openlibertywebprofile8_cmd_zip := exec.Command("curl", "-L", openlibertywebprofile8_zip_url, "-o", "package.zip")
	err = openlibertywebprofile8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openlibertywebprofile8_bin_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-webProfile8-24.0.0.8.zip"
	openlibertywebprofile8_cmd_bin := exec.Command("curl", "-L", openlibertywebprofile8_bin_url, "-o", "binary.bin")
	err = openlibertywebprofile8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openlibertywebprofile8_src_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/24.0.0.8/openliberty-webProfile8-24.0.0.8.zip"
	openlibertywebprofile8_cmd_src := exec.Command("curl", "-L", openlibertywebprofile8_src_url, "-o", "source.tar.gz")
	err = openlibertywebprofile8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openlibertywebprofile8_cmd_direct := exec.Command("./binary")
	err = openlibertywebprofile8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

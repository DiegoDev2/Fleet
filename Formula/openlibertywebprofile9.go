package main

// OpenlibertyWebprofile9 - Lightweight open framework for Java (Jakarta EE Web Profile 9)
// Homepage: https://openliberty.io

import (
	"fmt"
	
	"os/exec"
)

func installOpenlibertyWebprofile9() {
	// Método 1: Descargar y extraer .tar.gz
	openlibertywebprofile9_tar_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/23.0.0.2/openliberty-webProfile9-23.0.0.2.zip"
	openlibertywebprofile9_cmd_tar := exec.Command("curl", "-L", openlibertywebprofile9_tar_url, "-o", "package.tar.gz")
	err := openlibertywebprofile9_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openlibertywebprofile9_zip_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/23.0.0.2/openliberty-webProfile9-23.0.0.2.zip"
	openlibertywebprofile9_cmd_zip := exec.Command("curl", "-L", openlibertywebprofile9_zip_url, "-o", "package.zip")
	err = openlibertywebprofile9_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openlibertywebprofile9_bin_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/23.0.0.2/openliberty-webProfile9-23.0.0.2.zip"
	openlibertywebprofile9_cmd_bin := exec.Command("curl", "-L", openlibertywebprofile9_bin_url, "-o", "binary.bin")
	err = openlibertywebprofile9_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openlibertywebprofile9_src_url := "https://public.dhe.ibm.com/ibmdl/export/pub/software/openliberty/runtime/release/23.0.0.2/openliberty-webProfile9-23.0.0.2.zip"
	openlibertywebprofile9_cmd_src := exec.Command("curl", "-L", openlibertywebprofile9_src_url, "-o", "source.tar.gz")
	err = openlibertywebprofile9_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openlibertywebprofile9_cmd_direct := exec.Command("./binary")
	err = openlibertywebprofile9_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

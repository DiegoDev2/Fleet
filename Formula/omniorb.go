package main

// Omniorb - IOR and naming service utilities for omniORB
// Homepage: https://omniorb.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installOmniorb() {
	// Método 1: Descargar y extraer .tar.gz
	omniorb_tar_url := "https://downloads.sourceforge.net/project/omniorb/omniORB/omniORB-4.3.2/omniORB-4.3.2.tar.bz2"
	omniorb_cmd_tar := exec.Command("curl", "-L", omniorb_tar_url, "-o", "package.tar.gz")
	err := omniorb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	omniorb_zip_url := "https://downloads.sourceforge.net/project/omniorb/omniORB/omniORB-4.3.2/omniORB-4.3.2.tar.bz2"
	omniorb_cmd_zip := exec.Command("curl", "-L", omniorb_zip_url, "-o", "package.zip")
	err = omniorb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	omniorb_bin_url := "https://downloads.sourceforge.net/project/omniorb/omniORB/omniORB-4.3.2/omniORB-4.3.2.tar.bz2"
	omniorb_cmd_bin := exec.Command("curl", "-L", omniorb_bin_url, "-o", "binary.bin")
	err = omniorb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	omniorb_src_url := "https://downloads.sourceforge.net/project/omniorb/omniORB/omniORB-4.3.2/omniORB-4.3.2.tar.bz2"
	omniorb_cmd_src := exec.Command("curl", "-L", omniorb_src_url, "-o", "source.tar.gz")
	err = omniorb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	omniorb_cmd_direct := exec.Command("./binary")
	err = omniorb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}

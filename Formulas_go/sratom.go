package main

// Sratom - Library for serializing LV2 atoms to/from RDF
// Homepage: https://drobilla.net/software/sratom.html

import (
	"fmt"
	
	"os/exec"
)

func installSratom() {
	// Método 1: Descargar y extraer .tar.gz
	sratom_tar_url := "https://download.drobilla.net/sratom-0.6.16.tar.xz"
	sratom_cmd_tar := exec.Command("curl", "-L", sratom_tar_url, "-o", "package.tar.gz")
	err := sratom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sratom_zip_url := "https://download.drobilla.net/sratom-0.6.16.tar.xz"
	sratom_cmd_zip := exec.Command("curl", "-L", sratom_zip_url, "-o", "package.zip")
	err = sratom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sratom_bin_url := "https://download.drobilla.net/sratom-0.6.16.tar.xz"
	sratom_cmd_bin := exec.Command("curl", "-L", sratom_bin_url, "-o", "binary.bin")
	err = sratom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sratom_src_url := "https://download.drobilla.net/sratom-0.6.16.tar.xz"
	sratom_cmd_src := exec.Command("curl", "-L", sratom_src_url, "-o", "source.tar.gz")
	err = sratom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sratom_cmd_direct := exec.Command("./binary")
	err = sratom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: lv2")
exec.Command("latte", "install", "lv2")
	fmt.Println("Instalando dependencia: serd")
exec.Command("latte", "install", "serd")
	fmt.Println("Instalando dependencia: sord")
exec.Command("latte", "install", "sord")
}

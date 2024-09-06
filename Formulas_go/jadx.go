package main

// Jadx - Dex to Java decompiler
// Homepage: https://github.com/skylot/jadx

import (
	"fmt"
	
	"os/exec"
)

func installJadx() {
	// Método 1: Descargar y extraer .tar.gz
	jadx_tar_url := "https://github.com/skylot/jadx/releases/download/v1.5.0/jadx-1.5.0.zip"
	jadx_cmd_tar := exec.Command("curl", "-L", jadx_tar_url, "-o", "package.tar.gz")
	err := jadx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jadx_zip_url := "https://github.com/skylot/jadx/releases/download/v1.5.0/jadx-1.5.0.zip"
	jadx_cmd_zip := exec.Command("curl", "-L", jadx_zip_url, "-o", "package.zip")
	err = jadx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jadx_bin_url := "https://github.com/skylot/jadx/releases/download/v1.5.0/jadx-1.5.0.zip"
	jadx_cmd_bin := exec.Command("curl", "-L", jadx_bin_url, "-o", "binary.bin")
	err = jadx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jadx_src_url := "https://github.com/skylot/jadx/releases/download/v1.5.0/jadx-1.5.0.zip"
	jadx_cmd_src := exec.Command("curl", "-L", jadx_src_url, "-o", "source.tar.gz")
	err = jadx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jadx_cmd_direct := exec.Command("./binary")
	err = jadx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
exec.Command("latte", "install", "gradle")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}

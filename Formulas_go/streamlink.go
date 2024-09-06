package main

// Streamlink - CLI for extracting streams from various websites to a video player
// Homepage: https://streamlink.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installStreamlink() {
	// Método 1: Descargar y extraer .tar.gz
	streamlink_tar_url := "https://files.pythonhosted.org/packages/88/a9/1c15931891a4e225a12b54a60d6e3e643a51678d4e7253517ff8e8474703/streamlink-6.9.0.tar.gz"
	streamlink_cmd_tar := exec.Command("curl", "-L", streamlink_tar_url, "-o", "package.tar.gz")
	err := streamlink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	streamlink_zip_url := "https://files.pythonhosted.org/packages/88/a9/1c15931891a4e225a12b54a60d6e3e643a51678d4e7253517ff8e8474703/streamlink-6.9.0.zip"
	streamlink_cmd_zip := exec.Command("curl", "-L", streamlink_zip_url, "-o", "package.zip")
	err = streamlink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	streamlink_bin_url := "https://files.pythonhosted.org/packages/88/a9/1c15931891a4e225a12b54a60d6e3e643a51678d4e7253517ff8e8474703/streamlink-6.9.0.bin"
	streamlink_cmd_bin := exec.Command("curl", "-L", streamlink_bin_url, "-o", "binary.bin")
	err = streamlink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	streamlink_src_url := "https://files.pythonhosted.org/packages/88/a9/1c15931891a4e225a12b54a60d6e3e643a51678d4e7253517ff8e8474703/streamlink-6.9.0.src.tar.gz"
	streamlink_cmd_src := exec.Command("curl", "-L", streamlink_src_url, "-o", "source.tar.gz")
	err = streamlink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	streamlink_cmd_direct := exec.Command("./binary")
	err = streamlink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}

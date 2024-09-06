package main

// Hyfetch - Fast, highly customisable system info script with LGBTQ+ pride flags
// Homepage: https://github.com/hykilpikonna/hyfetch

import (
	"fmt"
	
	"os/exec"
)

func installHyfetch() {
	// Método 1: Descargar y extraer .tar.gz
	hyfetch_tar_url := "https://files.pythonhosted.org/packages/bb/af/0c4590b16c84073bd49b09ada0756fd9bd75b072e3ba9aec73101f0cc9f4/HyFetch-1.4.11.tar.gz"
	hyfetch_cmd_tar := exec.Command("curl", "-L", hyfetch_tar_url, "-o", "package.tar.gz")
	err := hyfetch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hyfetch_zip_url := "https://files.pythonhosted.org/packages/bb/af/0c4590b16c84073bd49b09ada0756fd9bd75b072e3ba9aec73101f0cc9f4/HyFetch-1.4.11.zip"
	hyfetch_cmd_zip := exec.Command("curl", "-L", hyfetch_zip_url, "-o", "package.zip")
	err = hyfetch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hyfetch_bin_url := "https://files.pythonhosted.org/packages/bb/af/0c4590b16c84073bd49b09ada0756fd9bd75b072e3ba9aec73101f0cc9f4/HyFetch-1.4.11.bin"
	hyfetch_cmd_bin := exec.Command("curl", "-L", hyfetch_bin_url, "-o", "binary.bin")
	err = hyfetch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hyfetch_src_url := "https://files.pythonhosted.org/packages/bb/af/0c4590b16c84073bd49b09ada0756fd9bd75b072e3ba9aec73101f0cc9f4/HyFetch-1.4.11.src.tar.gz"
	hyfetch_cmd_src := exec.Command("curl", "-L", hyfetch_src_url, "-o", "source.tar.gz")
	err = hyfetch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hyfetch_cmd_direct := exec.Command("./binary")
	err = hyfetch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}

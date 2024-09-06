package main

// OpenaiWhisper - General-purpose speech recognition model
// Homepage: https://github.com/openai/whisper

import (
	"fmt"
	
	"os/exec"
)

func installOpenaiWhisper() {
	// Método 1: Descargar y extraer .tar.gz
	openaiwhisper_tar_url := "https://files.pythonhosted.org/packages/d2/6e/50ace2bf704e5ffc786d20d96403ab0d57c5d6ab8729de7fed8c436687df/openai-whisper-20231117.tar.gz"
	openaiwhisper_cmd_tar := exec.Command("curl", "-L", openaiwhisper_tar_url, "-o", "package.tar.gz")
	err := openaiwhisper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openaiwhisper_zip_url := "https://files.pythonhosted.org/packages/d2/6e/50ace2bf704e5ffc786d20d96403ab0d57c5d6ab8729de7fed8c436687df/openai-whisper-20231117.zip"
	openaiwhisper_cmd_zip := exec.Command("curl", "-L", openaiwhisper_zip_url, "-o", "package.zip")
	err = openaiwhisper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openaiwhisper_bin_url := "https://files.pythonhosted.org/packages/d2/6e/50ace2bf704e5ffc786d20d96403ab0d57c5d6ab8729de7fed8c436687df/openai-whisper-20231117.bin"
	openaiwhisper_cmd_bin := exec.Command("curl", "-L", openaiwhisper_bin_url, "-o", "binary.bin")
	err = openaiwhisper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openaiwhisper_src_url := "https://files.pythonhosted.org/packages/d2/6e/50ace2bf704e5ffc786d20d96403ab0d57c5d6ab8729de7fed8c436687df/openai-whisper-20231117.src.tar.gz"
	openaiwhisper_cmd_src := exec.Command("curl", "-L", openaiwhisper_src_url, "-o", "source.tar.gz")
	err = openaiwhisper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openaiwhisper_cmd_direct := exec.Command("./binary")
	err = openaiwhisper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: llvm@15")
	exec.Command("latte", "install", "llvm@15").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pytorch")
	exec.Command("latte", "install", "pytorch").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}

package main

// VapoursynthSub - VapourSynth filters - Subtitling filter
// Homepage: https://www.vapoursynth.com

import (
	"fmt"
	
	"os/exec"
)

func installVapoursynthSub() {
	// Método 1: Descargar y extraer .tar.gz
	vapoursynthsub_tar_url := "https://github.com/vapoursynth/subtext/archive/refs/tags/R5.tar.gz"
	vapoursynthsub_cmd_tar := exec.Command("curl", "-L", vapoursynthsub_tar_url, "-o", "package.tar.gz")
	err := vapoursynthsub_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vapoursynthsub_zip_url := "https://github.com/vapoursynth/subtext/archive/refs/tags/R5.zip"
	vapoursynthsub_cmd_zip := exec.Command("curl", "-L", vapoursynthsub_zip_url, "-o", "package.zip")
	err = vapoursynthsub_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vapoursynthsub_bin_url := "https://github.com/vapoursynth/subtext/archive/refs/tags/R5.bin"
	vapoursynthsub_cmd_bin := exec.Command("curl", "-L", vapoursynthsub_bin_url, "-o", "binary.bin")
	err = vapoursynthsub_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vapoursynthsub_src_url := "https://github.com/vapoursynth/subtext/archive/refs/tags/R5.src.tar.gz"
	vapoursynthsub_cmd_src := exec.Command("curl", "-L", vapoursynthsub_src_url, "-o", "source.tar.gz")
	err = vapoursynthsub_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vapoursynthsub_cmd_direct := exec.Command("./binary")
	err = vapoursynthsub_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: libass")
exec.Command("latte", "install", "libass")
	fmt.Println("Instalando dependencia: vapoursynth")
exec.Command("latte", "install", "vapoursynth")
}

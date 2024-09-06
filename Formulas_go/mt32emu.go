package main

// Mt32emu - Multi-platform software synthesiser
// Homepage: https://github.com/munt/munt

import (
	"fmt"
	
	"os/exec"
)

func installMt32emu() {
	// Método 1: Descargar y extraer .tar.gz
	mt32emu_tar_url := "https://github.com/munt/munt/archive/refs/tags/libmt32emu_2_7_1.tar.gz"
	mt32emu_cmd_tar := exec.Command("curl", "-L", mt32emu_tar_url, "-o", "package.tar.gz")
	err := mt32emu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mt32emu_zip_url := "https://github.com/munt/munt/archive/refs/tags/libmt32emu_2_7_1.zip"
	mt32emu_cmd_zip := exec.Command("curl", "-L", mt32emu_zip_url, "-o", "package.zip")
	err = mt32emu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mt32emu_bin_url := "https://github.com/munt/munt/archive/refs/tags/libmt32emu_2_7_1.bin"
	mt32emu_cmd_bin := exec.Command("curl", "-L", mt32emu_bin_url, "-o", "binary.bin")
	err = mt32emu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mt32emu_src_url := "https://github.com/munt/munt/archive/refs/tags/libmt32emu_2_7_1.src.tar.gz"
	mt32emu_cmd_src := exec.Command("curl", "-L", mt32emu_src_url, "-o", "source.tar.gz")
	err = mt32emu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mt32emu_cmd_direct := exec.Command("./binary")
	err = mt32emu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsoxr")
exec.Command("latte", "install", "libsoxr")
}

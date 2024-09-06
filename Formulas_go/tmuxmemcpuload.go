package main

// TmuxMemCpuLoad - CPU, RAM memory, and load monitor for use with tmux
// Homepage: https://github.com/thewtex/tmux-mem-cpu-load

import (
	"fmt"
	
	"os/exec"
)

func installTmuxMemCpuLoad() {
	// Método 1: Descargar y extraer .tar.gz
	tmuxmemcpuload_tar_url := "https://github.com/thewtex/tmux-mem-cpu-load/archive/refs/tags/v3.8.1.tar.gz"
	tmuxmemcpuload_cmd_tar := exec.Command("curl", "-L", tmuxmemcpuload_tar_url, "-o", "package.tar.gz")
	err := tmuxmemcpuload_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tmuxmemcpuload_zip_url := "https://github.com/thewtex/tmux-mem-cpu-load/archive/refs/tags/v3.8.1.zip"
	tmuxmemcpuload_cmd_zip := exec.Command("curl", "-L", tmuxmemcpuload_zip_url, "-o", "package.zip")
	err = tmuxmemcpuload_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tmuxmemcpuload_bin_url := "https://github.com/thewtex/tmux-mem-cpu-load/archive/refs/tags/v3.8.1.bin"
	tmuxmemcpuload_cmd_bin := exec.Command("curl", "-L", tmuxmemcpuload_bin_url, "-o", "binary.bin")
	err = tmuxmemcpuload_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tmuxmemcpuload_src_url := "https://github.com/thewtex/tmux-mem-cpu-load/archive/refs/tags/v3.8.1.src.tar.gz"
	tmuxmemcpuload_cmd_src := exec.Command("curl", "-L", tmuxmemcpuload_src_url, "-o", "source.tar.gz")
	err = tmuxmemcpuload_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tmuxmemcpuload_cmd_direct := exec.Command("./binary")
	err = tmuxmemcpuload_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}

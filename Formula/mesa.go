package main

// Mesa - Graphics Library
// Homepage: https://www.mesa3d.org/

import (
	"fmt"
	
	"os/exec"
)

func installMesa() {
	// Método 1: Descargar y extraer .tar.gz
	mesa_tar_url := "https://mesa.freedesktop.org/archive/mesa-24.2.1.tar.xz"
	mesa_cmd_tar := exec.Command("curl", "-L", mesa_tar_url, "-o", "package.tar.gz")
	err := mesa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mesa_zip_url := "https://mesa.freedesktop.org/archive/mesa-24.2.1.tar.xz"
	mesa_cmd_zip := exec.Command("curl", "-L", mesa_zip_url, "-o", "package.zip")
	err = mesa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mesa_bin_url := "https://mesa.freedesktop.org/archive/mesa-24.2.1.tar.xz"
	mesa_cmd_bin := exec.Command("curl", "-L", mesa_bin_url, "-o", "binary.bin")
	err = mesa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mesa_src_url := "https://mesa.freedesktop.org/archive/mesa-24.2.1.tar.xz"
	mesa_cmd_src := exec.Command("curl", "-L", mesa_src_url, "-o", "source.tar.gz")
	err = mesa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mesa_cmd_direct := exec.Command("./binary")
	err = mesa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: xorgproto")
	exec.Command("latte", "install", "xorgproto").Run()
	fmt.Println("Instalando dependencia: expat")
	exec.Command("latte", "install", "expat").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxfixes")
	exec.Command("latte", "install", "libxfixes").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: glslang")
	exec.Command("latte", "install", "glslang").Run()
	fmt.Println("Instalando dependencia: gzip")
	exec.Command("latte", "install", "gzip").Run()
	fmt.Println("Instalando dependencia: libclc")
	exec.Command("latte", "install", "libclc").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libva")
	exec.Command("latte", "install", "libva").Run()
	fmt.Println("Instalando dependencia: libvdpau")
	exec.Command("latte", "install", "libvdpau").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: libxshmfence")
	exec.Command("latte", "install", "libxshmfence").Run()
	fmt.Println("Instalando dependencia: libxv")
	exec.Command("latte", "install", "libxv").Run()
	fmt.Println("Instalando dependencia: libxxf86vm")
	exec.Command("latte", "install", "libxxf86vm").Run()
	fmt.Println("Instalando dependencia: lm-sensors")
	exec.Command("latte", "install", "lm-sensors").Run()
	fmt.Println("Instalando dependencia: spirv-llvm-translator")
	exec.Command("latte", "install", "spirv-llvm-translator").Run()
	fmt.Println("Instalando dependencia: valgrind")
	exec.Command("latte", "install", "valgrind").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
	fmt.Println("Instalando dependencia: wayland-protocols")
	exec.Command("latte", "install", "wayland-protocols").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}

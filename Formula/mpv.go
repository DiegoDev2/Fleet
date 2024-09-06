package main

// Mpv - Media player based on MPlayer and mplayer2
// Homepage: https://mpv.io

import (
	"fmt"
	
	"os/exec"
)

func installMpv() {
	// Método 1: Descargar y extraer .tar.gz
	mpv_tar_url := "https://github.com/mpv-player/mpv/archive/refs/tags/v0.38.0.tar.gz"
	mpv_cmd_tar := exec.Command("curl", "-L", mpv_tar_url, "-o", "package.tar.gz")
	err := mpv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpv_zip_url := "https://github.com/mpv-player/mpv/archive/refs/tags/v0.38.0.zip"
	mpv_cmd_zip := exec.Command("curl", "-L", mpv_zip_url, "-o", "package.zip")
	err = mpv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpv_bin_url := "https://github.com/mpv-player/mpv/archive/refs/tags/v0.38.0.bin"
	mpv_cmd_bin := exec.Command("curl", "-L", mpv_bin_url, "-o", "binary.bin")
	err = mpv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpv_src_url := "https://github.com/mpv-player/mpv/archive/refs/tags/v0.38.0.src.tar.gz"
	mpv_cmd_src := exec.Command("curl", "-L", mpv_src_url, "-o", "source.tar.gz")
	err = mpv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpv_cmd_direct := exec.Command("./binary")
	err = mpv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docutils")
	exec.Command("latte", "install", "docutils").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libass")
	exec.Command("latte", "install", "libass").Run()
	fmt.Println("Instalando dependencia: libbluray")
	exec.Command("latte", "install", "libbluray").Run()
	fmt.Println("Instalando dependencia: libplacebo")
	exec.Command("latte", "install", "libplacebo").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: mujs")
	exec.Command("latte", "install", "mujs").Run()
	fmt.Println("Instalando dependencia: rubberband")
	exec.Command("latte", "install", "rubberband").Run()
	fmt.Println("Instalando dependencia: uchardet")
	exec.Command("latte", "install", "uchardet").Run()
	fmt.Println("Instalando dependencia: vapoursynth")
	exec.Command("latte", "install", "vapoursynth").Run()
	fmt.Println("Instalando dependencia: vulkan-loader")
	exec.Command("latte", "install", "vulkan-loader").Run()
	fmt.Println("Instalando dependencia: yt-dlp")
	exec.Command("latte", "install", "yt-dlp").Run()
	fmt.Println("Instalando dependencia: zimg")
	exec.Command("latte", "install", "zimg").Run()
	fmt.Println("Instalando dependencia: molten-vk")
	exec.Command("latte", "install", "molten-vk").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: libdrm")
	exec.Command("latte", "install", "libdrm").Run()
	fmt.Println("Instalando dependencia: libva")
	exec.Command("latte", "install", "libva").Run()
	fmt.Println("Instalando dependencia: libvdpau")
	exec.Command("latte", "install", "libvdpau").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxkbcommon")
	exec.Command("latte", "install", "libxkbcommon").Run()
	fmt.Println("Instalando dependencia: libxpresent")
	exec.Command("latte", "install", "libxpresent").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: libxscrnsaver")
	exec.Command("latte", "install", "libxscrnsaver").Run()
	fmt.Println("Instalando dependencia: libxv")
	exec.Command("latte", "install", "libxv").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
	fmt.Println("Instalando dependencia: wayland")
	exec.Command("latte", "install", "wayland").Run()
}
